package message

import "github.com/nilorg/go-wechat/v2/client"

// Message 消息
type Message struct {
	client client.Clienter
}

// NewMessage ...
func NewMessage(c client.Clienter) *Message {
	return &Message{
		client: c,
	}
}
