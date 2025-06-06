package chat

import "time"

type Message struct {
	Content string
	Author  Author
	Time    time.Time
}

func NewMessage(content string, author Author) *Message {
	return &Message{
		Content: content,
		Author:  author,
		Time:    time.Now(),
	}
}
