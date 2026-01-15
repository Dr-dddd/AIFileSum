package handler

import (
	"AIFileSum/service"
	"io"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file is required"})
		return
	}

	if filepath.Ext(file.Filename) != ".txt" {
		c.JSON(400, gin.H{"error": "only .txt file supported"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot open file"})
		return
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot read file"})
		return
	}

	text := string(content)

	// TODO：调用大模型总结
	summary := service.MockSummary(text)

	c.JSON(200, gin.H{
		"summary": summary,
	})
}
