package models

type QwenRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type QwenResponse struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
