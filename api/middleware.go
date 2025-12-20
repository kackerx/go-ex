package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyDumpResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *bodyDumpResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

type CachedResponse struct {
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

// === Idenpotency
const (
	HeaderIdempotencyKey = "X-Idempotency-Key"
	KeyPrefix            = "idempotency:"
	LockExpire           = 24 * time.Hour
)

func IdempotencyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idempotencyKey := c.GetHeader(HeaderIdempotencyKey)
		if idempotencyKey == "" {
			c.Next()
			return
		}
		cacheKey := KeyPrefix + idempotencyKey
		val, err := redisClient.Get(c.Request.Context(), cacheKey).Result()
		if err == nil {
			if val == "PROCESSING" {
				c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
				c.Abort()
				return
			}

			var cachedResponse CachedResponse
			json.Unmarshal([]byte(val), &cachedResponse)
			for k, v := range cachedResponse.Headers {
				for _, vv := range v {
					c.Header(k, vv)
				}
			}
			c.Data(cachedResponse.Status, c.Writer.Header().Get("Content-Type"), []byte(cachedResponse.Body))
			c.Abort()
			return
		}

		success, _ := redisClient.SetNX(c.Request.Context(), cacheKey, "PROCESSING", LockExpire).Result()
		if !success {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			c.Abort()
			return
		}

		dumpWriter := &bodyDumpResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = dumpWriter

		c.Next()

		if c.Writer.Status() < 500 {
			cacheResponse := CachedResponse{
				Status:  c.Writer.Status(),
				Headers: c.Writer.Header(),
				Body:    dumpWriter.body.String(),
			}
			jsonResponse, _ := json.Marshal(cacheResponse)
			redisClient.Set(c.Request.Context(), cacheKey, string(jsonResponse), LockExpire)
		} else {
			redisClient.Del(c.Request.Context(), cacheKey)
		}
	}
}
