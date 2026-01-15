package service

func MockSummary(text string) string {
	if len(text) > 200 {
		text = text[:200]
	}
	return "【模拟总结】文章主要内容是：" + text
}
