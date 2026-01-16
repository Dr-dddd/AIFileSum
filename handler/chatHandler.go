package handler

import (
	"AIFileSum/models"
	"AIFileSum/service"
	"io"

	"github.com/gin-gonic/gin"
)

func ChatHandler(c *gin.Context) {

	msgType := c.PostForm("type")
	if msgType != "file" {
		c.JSON(400, gin.H{"error": "only file supported now"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file required"})
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

	summary, err := service.CallQwenSummary(text)
	if err != nil {
		c.JSON(500, gin.H{"error": "call qwen error"})
		return
	}

	c.JSON(200, models.Message{
		Role:    "assistant",
		Type:    "text",
		Content: summary,
	})
}
