package client

import (
	"encoding/json"
	"fmt"
)

// UrlSchemeGenerateRequest 生成小程序码
type UrlSchemeGenerateRequest struct {
	JumpWxa        *UrlSchemeJumpWxa `json:"jump_wxa"`        // 跳转到的目标小程序信息。
	ExpireType     int               `json:"expire_type"`     // 到期失效的 scheme 码失效类型，失效时间：0，失效间隔天数：1
	ExpireTime     int               `json:"expire_time"`     // 到期失效的 scheme 码的失效时间，为 Unix 时间戳。生成的到期失效 scheme 码在该时间前有效。最长有效期为30天。expire_type 为 0 时必填
	ExpireInterval int               `json:"expire_interval"` // 到期失效的 scheme 码的失效间隔天数。生成的到期失效 scheme 码在该间隔时间到达前有效。最长间隔天数为30天。 expire_type 为 1 时必填
}

type UrlSchemeJumpWxa struct {
	Path       string `json:"path"`        // 通过 scheme 码进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带 query。path 为空时会跳转小程序主页。
	Query      string `json:"query"`       // 通过 scheme 码进入小程序时的 query，最大1024个字符，只支持数字，大小写英文以及部分特殊字符：`!#$&'()*+,/:;=?@-._~%``
	EnvVersion string `json:"env_version"` // 要打开的小程序版本。正式版为"release"，体验版为"trial"，开发版为"develop"，仅在微信外打开时生效。
}

// UrlSchemeGenerateResponse 生成小程序码
type UrlSchemeGenerateResponse struct {
	Openlink string `json:"openlink"` // 生成的 scheme 码
}

// UrlSchemeGenerate 生成小程序码
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html
func (c *Client) UrlSchemeGenerate(req *UrlSchemeGenerateRequest) (*UrlSchemeGenerateResponse, error) {
	url := fmt.Sprintf("%s/wxa/generatescheme", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	reply := new(UrlSchemeGenerateResponse)
	err = json.Unmarshal(result, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
