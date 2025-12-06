package main

import "github.com/gin-gonic/gin"

func Hehe(c *gin.Context) {
	c.Set("biz", "Notify")
	c.JSON(200, gin.H{"biz": "Notify"})
}
