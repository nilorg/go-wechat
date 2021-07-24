package client

type Client2 struct {
	opts ClientOptions
}

// ClientOptions 可选参数列表
type ClientOptions struct {
	BaseURL   string
	Proxy     bool
	Token     Tokener
	AppID     string
	AppSecret string
}

// ClientOption 为可选参数赋值的函数
type ClientOption func(*ClientOptions)

// NewClient2Options 创建可选参数
func NewClient2Options(opts ...ClientOption) ClientOptions {
	opt := ClientOptions{
		BaseURL: "https://api.weixin.qq.com",
		Proxy:   false,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// ClientOptionBaseURL ...
func ClientOptionBaseURL(baseURL string) ClientOption {
	return func(o *ClientOptions) {
		o.BaseURL = baseURL
	}
}

// ClientOptionProxy ...
func ClientOptionProxy(proxy bool) ClientOption {
	return func(o *ClientOptions) {
		o.Proxy = proxy
	}
}

// ClientOptionToken ...
func ClientOptionToken(token Tokener) ClientOption {
	return func(o *ClientOptions) {
		o.Token = token
	}
}

func NewClient2(opts ...ClientOption) (client *Client2) {
	client = new(Client2)
	client.opts = NewClient2Options(opts...)
	return
}
