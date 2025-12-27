package main

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/kackerx/kai/app/interfaces/api"
	"github.com/redis/go-redis/v9"
)

type LimiterParams struct {
	Key    string        // 限流key
	Rate   int           // 周期内速率
	Period time.Duration // 周期
	Burst  int           // 突发量
}

type Limiter struct {
	core *redis_rate.Limiter
}

func NewLimiter(redis *redis.Client) *Limiter {
	return &Limiter{
		core: redis_rate.NewLimiter(redis),
	}
}

func (l *Limiter) Allow(ctx context.Context, key KeyStrategy, params LimiterParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		allow, err := l.core.Allow(ctx, key(c), redis_rate.Limit{
			Rate:   params.Rate,
			Burst:  params.Burst,
			Period: params.Period,
		})
		if err != nil {
			slog.Error("limiter allow error", "error", err)
			c.Next()
			return
		}

		c.Header("X-RateLimit-Limit", strconv.Itoa(allow.Limit.Burst))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(allow.Remaining))
		c.Header("X-RateLimit-Reset", strconv.Itoa(int(allow.ResetAfter.Seconds())))

		if allow.Allowed == 0 {
			retryAfter := allow.RetryAfter.Seconds()
			c.Header("Retry-After", strconv.Itoa(int(retryAfter)))
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}

type KeyStrategy func(c *gin.Context) string

func IPKeyStrategy(c *gin.Context) string {
	return "ip:" + c.ClientIP()
}

func UserIDKeyStrategy(c *gin.Context) string {
	return "user_id:" + api.GetUserID(c)
}

func PathKeyStrategy(c *gin.Context) string {
	return "path:" + c.Request.URL.Path
}
