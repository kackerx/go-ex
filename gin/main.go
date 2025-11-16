package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	r.GET("/home/*.html", func(c *gin.Context) {
		page := c.Param(".html")
		slog.Info("page", "page", page)
		slog.Error("page", "page", page)
		slog.Warn("page", "page", page)
		slog.Info("todo", "TODO", page)
		slog.Info("select", "SELECT * FROM users WHERE id = 1", "")
		log.Println("df")
		fmt.Println("INFO SELECT")

		c.String(200, page)
	})

	r.Run(":8080")
}
