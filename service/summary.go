package service

import (
	"AIFileSum/config"
	"AIFileSum/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func CallQwenSummary(text string) (string, error) {
	cfg := config.Get()
	apiKey := cfg.Qwen.APIKey
	//if apiKey == "" {
	//	return "", errors.New("DASHSCOPE_API_KEY not set")
	//}

	reqBody := models.QwenRequest{
		Model: "qwen-plus",
		Messages: []models.Message{
			{
				Role:    "system",
				Content: "你是一个专业的文件总结助手，请用简洁的语言总结文件的核心内容",
			},
			{
				Role:    "user",
				Content: text,
			},
		},
	}

	data, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", cfg.Qwen.BaseURL, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res models.QwenResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	if len(res.Choices) == 0 {
		return "", errors.New("empty choices from qwen")
	}

	return res.Choices[0].Message.Content, nil
}
