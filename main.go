package main

import (
	"AIFileSum/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", handler.UploadHandler)

	r.Run(":8080")
}
