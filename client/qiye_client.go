package client

import "net/http"

type QiyeClient struct {
	opts QiyeClientOptions
}

// QiyeClientOptions 可选参数列表
type QiyeClientOptions struct {
	BaseURL    string
	Proxy      bool
	Token      QiyeTokener
	AppID      string
	AppSecret  string
	HttpClient *http.Client
}

// QiyeClientOption 为可选参数赋值的函数
type QiyeClientOption func(*QiyeClientOptions)

// NewQiyeClientOptions 创建可选参数
func NewQiyeClientOptions(opts ...QiyeClientOption) QiyeClientOptions {
	opt := QiyeClientOptions{
		BaseURL:    "https://qyapi.weixin.qq.com",
		Proxy:      false,
		HttpClient: http.DefaultClient,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// QiyeClientOptionBaseURL ...
func QiyeClientOptionBaseURL(baseURL string) QiyeClientOption {
	return func(o *QiyeClientOptions) {
		o.BaseURL = baseURL
	}
}

// QiyeClientOptionProxy ...
func QiyeClientOptionProxy(proxy bool) QiyeClientOption {
	return func(o *QiyeClientOptions) {
		o.Proxy = proxy
	}
}

// QiyeClientOptionToken ...
func QiyeClientOptionToken(token QiyeTokener) QiyeClientOption {
	return func(o *QiyeClientOptions) {
		o.Token = token
	}
}

// QiyeClientOptionAppID ...
func QiyeClientOptionAppID(appID string) QiyeClientOption {
	return func(o *QiyeClientOptions) {
		o.AppID = appID
	}
}

// QiyeClientOptionAppSecret ...
func QiyeClientOptionAppSecret(appSecret string) QiyeClientOption {
	return func(o *QiyeClientOptions) {
		o.AppSecret = appSecret
	}
}

// QiyeClientOptionHttpClient ...
func QiyeClientOptionHttpClient(httpClient *http.Client) QiyeClientOption {
	return func(o *QiyeClientOptions) {
		o.HttpClient = httpClient
	}
}

func NewQiyeClient(opts ...QiyeClientOption) (client *QiyeClient) {
	client = new(QiyeClient)
	client.opts = NewQiyeClientOptions(opts...)
	return
}
