package service

import (
	"AIFileSum/config"
	"AIFileSum/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

func UploadFileToQwen(filename string, data []byte) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// 1️⃣ 写 file 字段
	filePart, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return "", err
	}
	_, err = filePart.Write(data)
	if err != nil {
		return "", err
	}

	// 2️⃣ 写 purpose 字段（非常关键）
	err = writer.WriteField("purpose", "file-extract")
	if err != nil {
		return "", err
	}

	// 3️⃣ 关闭 writer（必须）
	err = writer.Close()
	if err != nil {
		return "", err
	}

	// 4️⃣ 构造请求
	req, err := http.NewRequest(
		"POST",
		"https://dashscope.aliyuncs.com/compatible-mode/v1/files",
		body,
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+config.Get().Qwen.APIKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 5️⃣ 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("upload failed: %s", string(b))
	}

	// 6️⃣ 解析返回
	var result struct {
		ID string `json:"id"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.ID, nil
}

func CallQwenSummary(fileID string) (string, error) {

	reqBody := models.QwenRequest{
		Model: "qwen-long",
		Messages: []models.Message{
			{
				Role:    "system",
				Content: "你是一个专业的文件总结助手，请用简洁的语言总结文件的核心内容",
			},
			{
				Role:    "system",
				Content: fmt.Sprintf("fileid://%s", fileID),
			},
			{
				Role:    "user",
				Content: "请总结这份文件的核心内容",
			},
		},
	}

	data, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", config.Get().Qwen.BaseURL, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Get().Qwen.APIKey)

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
