package main

import (
	"AIFileSum/config"
	"AIFileSum/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// dev 环境下处理跨域问题
	if config.Get().Env == "dev" {
		r.Use(func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}
			c.Next()
		})
	}

	r.POST("/chat", handler.ChatHandler)

	r.Run(":8080")
}
