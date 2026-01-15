package models

type Message struct {
	Role    string `json:"role"`    // user / assistant
	Type    string `json:"type"`    // text / file
	Content string `json:"content"` // 文本内容（文件时是文件内容）
}
