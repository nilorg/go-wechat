package message

import (
	wechat "github.com/nilorg/go-wechat"
)

// Message 消息
type Message struct {
	client *wechat.Client
}

// NewMessage ...
func NewMessage(c *wechat.Client) *Message {
	return &Message{
		client: c,
	}
}
