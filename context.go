package wechat

import "context"

type wechatKey struct{}

// FromContext 从上下文中获取微信客户端
func FromContext(ctx context.Context) (*Client, bool) {
	c, ok := ctx.Value(wechatKey{}).(*Client)
	return c, ok
}

// NewContext 创建微信客户端上下文
func NewContext(ctx context.Context, c *Client) context.Context {
	return context.WithValue(ctx, wechatKey{}, c)
}
