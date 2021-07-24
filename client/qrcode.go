package client

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// QrcodeTempQrcodeRequest ，临时二维码
type QrcodeTempQrcodeRequest struct {
	ExpireSeconds int                    `json:"expire_seconds"` // 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天），此字段如果不填，则默认有效期为30秒。
	ActionName    string                 `json:"action_name"`    // 二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值
	ActionInfo    map[string]interface{} `json:"action_info"`    // 二维码详细信息
}

// QrcodeTempQrcodeReply 临时二维码回复
type QrcodeTempQrcodeReply struct {
	Ticket        string `json:"ticket"`         // 获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。
	ExpireSeconds int    `json:"expire_seconds"` // 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天）。
	URL           string `json:"url"`            // 二维码图片解析后的地址，开发者可根据该地址自行生成需要的二维码图片
}

// QrcodeLimitQrcodeRequest 永久二维码
type QrcodeLimitQrcodeRequest struct {
	ActionName string                 `json:"action_name"` // 二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值
	ActionInfo map[string]interface{} `json:"action_info"` // 二维码详细信息
}

// QrcodeLimitQrcodeReply 永久二维码回复
type QrcodeLimitQrcodeReply struct {
	Ticket string `json:"ticket"` // 获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。
	URL    string `json:"url"`    // 二维码图片解析后的地址，开发者可根据该地址自行生成需要的二维码图片
}

// NewQrcodeTempStrQrcodeRequest 创建一个临时字符串二维码
// 场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
func NewQrcodeTempStrQrcodeRequest(str string, expireSeconds int) *QrcodeTempQrcodeRequest {
	if expireSeconds <= 0 {
		expireSeconds = 30
	}
	return &QrcodeTempQrcodeRequest{
		ExpireSeconds: expireSeconds,
		ActionName:    "QR_STR_SCENE", // 临时的字符串参数值
		ActionInfo: map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": str,
			},
		},
	}
}

// NewQrcodeTempQrcodeRequest 创建一个临时二维码
// 场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000）
func NewQrcodeTempQrcodeRequest(num uint, expireSeconds int) *QrcodeTempQrcodeRequest {
	if expireSeconds <= 0 {
		expireSeconds = 30
	}
	if num == 0 || num > 100000 {
		panic("临时二维码时为32位非0整型，永久二维码时最大值为100000")
	}
	return &QrcodeTempQrcodeRequest{
		ExpireSeconds: expireSeconds,
		ActionName:    "QR_SCENE", // 临时的整型参数值
		ActionInfo: map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_id": num,
			},
		},
	}
}

// NewQrcodeStrLimitQrcodeRequest 创建一个永久字符串二维码
// 场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
func NewQrcodeStrLimitQrcodeRequest(str string) *QrcodeLimitQrcodeRequest {
	l := len(str)
	if l == 0 || l > 64 {
		panic("长度限制为1到64")
	}
	return &QrcodeLimitQrcodeRequest{
		ActionName: "QR_LIMIT_STR_SCENE", // 永久的字符串参数值
		ActionInfo: map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": str,
			},
		},
	}
}

// NewQrcodeLimitQrcodeRequest 创建一个永久二维码
// 场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000）
func NewQrcodeLimitQrcodeRequest(num uint) *QrcodeLimitQrcodeRequest {
	if num == 0 || num > 100000 {
		panic("临时二维码时为32位非0整型，永久二维码时最大值为100000")
	}
	return &QrcodeLimitQrcodeRequest{
		ActionName: "QR_LIMIT_SCENE", // 永久的整型参数值
		ActionInfo: map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_id": num,
			},
		},
	}
}

// QrcodeCreateTemp 生成临时带参数的二维码
func (c *Client2) QrcodeCreateTemp(req *QrcodeTempQrcodeRequest) (*QrcodeTempQrcodeReply, error) {
	url := fmt.Sprintf("%s/cgi-bin/qrcode/create", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	reply := new(QrcodeTempQrcodeReply)
	json.Unmarshal(result, reply)
	return reply, nil
}

// QrcodeDownload 下载二维码
func (c *Client2) QrcodeDownload(ticket string) ([]byte, error) {
	ticket = url.QueryEscape(ticket)
	baseURL := "https://mp.weixin.qq.com"
	if c.opts.Proxy {
		baseURL = c.opts.BaseURL
	}
	result, err := Get(fmt.Sprintf("%s/cgi-bin/showqrcode", baseURL), map[string]string{
		"ticket": ticket,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QrcodeCreateLimit 生成一个永久带参数的二维码
func (c *Client2) QrcodeCreateLimit(req *QrcodeLimitQrcodeRequest) (*QrcodeLimitQrcodeReply, error) {
	url := fmt.Sprintf("%s/cgi-bin/qrcode/create", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	reply := new(QrcodeLimitQrcodeReply)
	json.Unmarshal(result, reply)
	return reply, nil
}
