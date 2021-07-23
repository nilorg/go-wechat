package contexts

import (
	"context"
	"errors"

	"github.com/nilorg/go-wechat/v2/client"
)

var (
	// ErrContextNotFoundClient 上下文不存在客户端错误
	ErrContextNotFoundClient = errors.New("上下文中没有获取到微信客户端")
)

type wechatKey struct{}

// FromContext 从上下文中获取微信客户端
func FromContext(ctx context.Context) (client.Clienter, error) {
	c, ok := ctx.Value(wechatKey{}).(client.Clienter)
	if !ok {
		return nil, ErrContextNotFoundClient
	}
	return c, nil
}

// NewContext 创建微信客户端上下文
func NewContext(ctx context.Context, c client.Clienter) context.Context {
	return context.WithValue(ctx, wechatKey{}, c)
}
