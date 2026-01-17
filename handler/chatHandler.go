package handler

import (
	"AIFileSum/models"
	"AIFileSum/service"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func ChatHandler(c *gin.Context) {
	// check type
	msgType := c.PostForm("type")
	if msgType != "file" {
		c.JSON(400, gin.H{"error": "only file supported now"})
		return
	}

	// read file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file required"})
		return
	}
	if file.Size > 150<<20 {
		c.JSON(400, gin.H{"error": "file too large"})
	}
	f, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot open file"})
		return
	}
	defer f.Close()
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot read file"})
		return
	}

	// upload
	fileID, err := service.UploadFileToQwen(file.Filename, fileBytes)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "upload file error"})
	}

	// Call LLM
	summary, err := service.CallQwenSummary(fileID)
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
