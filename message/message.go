package message

import (
	"github.com/nilorg/go-wechat/context"
)

// Message 消息
type Message struct {
	context *context.Context
}

// NewMessage ...
func NewMessage(c *context.Context) *Message {
	return &Message{
		context: c,
	}
}
