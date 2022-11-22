package client

import "net/http"

type Client struct {
	opts ClientOptions
}

// ClientOptions 可选参数列表
type ClientOptions struct {
	BaseURL    string
	Proxy      bool
	Token      Tokener
	AppID      string
	AppSecret  string
	HttpClient *http.Client
}

// ClientOption 为可选参数赋值的函数
type ClientOption func(*ClientOptions)

// NewClientOptions 创建可选参数
func NewClientOptions(opts ...ClientOption) ClientOptions {
	opt := ClientOptions{
		BaseURL:    "https://api.weixin.qq.com",
		Proxy:      false,
		HttpClient: http.DefaultClient,
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

// ClientOptionAppID ...
func ClientOptionAppID(appID string) ClientOption {
	return func(o *ClientOptions) {
		o.AppID = appID
	}
}

// ClientOptionAppSecret ...
func ClientOptionAppSecret(appSecret string) ClientOption {
	return func(o *ClientOptions) {
		o.AppSecret = appSecret
	}
}

// ClientOptionHttpClient ...
func ClientOptionHttpClient(httpClient *http.Client) ClientOption {
	return func(o *ClientOptions) {
		o.HttpClient = httpClient
	}
}

func NewClient(opts ...ClientOption) (client *Client) {
	client = new(Client)
	client.opts = NewClientOptions(opts...)
	return
}
