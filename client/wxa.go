package client

import (
	"encoding/json"
	"fmt"
)

// WxaGenerateUrlSchemeRequest 生成小程序码
type WxaGenerateUrlSchemeRequest struct {
	JumpWxa        *WxaGenerateUrlSchemeJumpWxa `json:"jump_wxa"`        // 跳转到的目标小程序信息。
	ExpireType     int                          `json:"expire_type"`     // 到期失效的 scheme 码失效类型，失效时间：0，失效间隔天数：1
	ExpireTime     int                          `json:"expire_time"`     // 到期失效的 scheme 码的失效时间，为 Unix 时间戳。生成的到期失效 scheme 码在该时间前有效。最长有效期为30天。expire_type 为 0 时必填
	ExpireInterval int                          `json:"expire_interval"` // 到期失效的 scheme 码的失效间隔天数。生成的到期失效 scheme 码在该间隔时间到达前有效。最长间隔天数为30天。 expire_type 为 1 时必填
}

type WxaGenerateUrlSchemeJumpWxa struct {
	Path       string `json:"path"`        // 通过 scheme 码进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带 query。path 为空时会跳转小程序主页。
	Query      string `json:"query"`       // 通过 scheme 码进入小程序时的 query，最大1024个字符，只支持数字，大小写英文以及部分特殊字符：`!#$&'()*+,/:;=?@-._~%``
	EnvVersion string `json:"env_version"` // 要打开的小程序版本。正式版为"release"，体验版为"trial"，开发版为"develop"，仅在微信外打开时生效。
}

// WxaGenerateUrlSchemeResponse 生成小程序码
type WxaGenerateUrlSchemeResponse struct {
	Openlink string `json:"openlink"` // 生成的 scheme 码
}

// WxaGenerateUrlScheme 生成小程序码
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-scheme/generateScheme.html
func (c *Client) WxaGenerateUrlScheme(req *WxaGenerateUrlSchemeRequest) (*WxaGenerateUrlSchemeResponse, error) {
	url := fmt.Sprintf("%s/wxa/generatescheme", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	reply := new(WxaGenerateUrlSchemeResponse)
	err = json.Unmarshal(result, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// WxaGenerateUrlLinkRequest 生成小程序码
type WxaGenerateUrlLinkRequest struct {
	Path           string                       `json:"path"`            // 通过 URL Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带 query 。path 为空时会跳转小程序主页
	Query          string                       `json:"query"`           // 通过 URL Link 进入小程序时的query，最大1024个字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~%
	IsExpire       bool                         `json:"is_expire"`       // 默认值false。生成的 URL Link 类型，到期失效：true，永久有效：false。注意，永久有效 Link 和有效时间超过180天的到期失效 Link 的总数上限为10万个，详见获取 URL Link，生成 Link 前请仔细确认。
	ExpireType     int                          `json:"expire_type"`     // 默认值0.小程序 URL Link 失效类型，失效时间：0，失效间隔天数：1
	ExpireTime     int64                        `json:"expire_time"`     // 到期失效的 URL Link 的失效时间，为 Unix 时间戳。生成的到期失效 URL Link 在该时间前有效。最长有效期为1年。expire_type 为 0 必填
	ExpireInterval int                          `json:"expire_interval"` // 到期失效的URL Link的失效间隔天数。生成的到期失效URL Link在该间隔时间到达前有效。最长间隔天数为365天。expire_type 为 1 必填
	CloudBase      *WxaGenerateUrlLinkCloudBase `json:"cloud_base"`      // 云开发静态网站自定义 H5 配置参数，可配置中转的云开发 H5 页面。不填默认用官方 H5 页面
	EnvVersion     string                       `json:"env_version"`     // 默认值"release"。要打开的小程序版本。正式版为 "release"，体验版为"trial"，开发版为"develop"，仅在微信外打开时生效
}

type WxaGenerateUrlLinkCloudBase struct {
	Env           string `json:"env"`            // 云开发环境
	Domain        string `json:"domain"`         // 静态网站自定义域名，不填则使用默认域名
	Path          string `json:"path"`           // 云开发静态网站 H5 页面路径，不可携带 query
	Query         string `json:"query"`          // 云开发静态网站 H5 页面 query 参数，最大 1024 个字符，只支持数字，大小写英文以及部分特殊字符：`!#$&'()*+,/:;=?@-._~%``
	ResourceAppID string `json:"resource_appid"` // 第三方批量代云开发时必填，表示创建该 env 的 appid （小程序/第三方平台）
}

// WxaGenerateUrlLinkResponse 生成小程序码
type WxaGenerateUrlLinkResponse struct {
	URLLink string `json:"url_link"` // 生成的 url link 码
}

// WxaGenerateUrlLink 生成小程序码
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-link/generateUrlLink.html
func (c *Client) WxaGenerateUrlLink(req *WxaGenerateUrlLinkRequest) (*WxaGenerateUrlLinkResponse, error) {
	url := fmt.Sprintf("%s/wxa/generate_urllink", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	reply := new(WxaGenerateUrlLinkResponse)
	err = json.Unmarshal(result, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
