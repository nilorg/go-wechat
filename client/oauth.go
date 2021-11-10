package client

import (
	"encoding/json"
	"fmt"

	"github.com/nilorg/go-wechat/v2/client/lang"
)

// OAuthAccessTokenReply 获取Token
type OAuthAccessTokenReply struct {
	AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int    `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
	OpenID       string `json:"openid"`        // 用户唯一标识，请注意，在未关注公众号时，用户访问公众号的网页，也会产生一个用户和公众号唯一的OpenID
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
}

// OAuthRefreshTokenReply 刷新Token
type OAuthRefreshTokenReply struct {
	AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int    `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
	OpenID       string `json:"openid"`        // 用户唯一标识
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
}

// OAuthUserInfoReply 用户信息
type OAuthUserInfoReply struct {
	OpenID     string   `json:"openid"`     // 用户的唯一标识
	NickName   string   `json:"nickname"`   // 用户昵称
	Sex        string   `json:"sex"`        // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Province   string   `json:"province"`   // 用户个人资料填写的省份
	City       string   `json:"city"`       // 普通用户个人资料填写的城市
	Country    string   `json:"country"`    // 国家，如中国为CN
	HeadimgURL string   `json:"headimgurl"` // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Privilege  []string `json:"privilege"`  // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	UnionID    string   `json:"unionid"`    // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
}

// OAuthCode2SessionReponse 小程序验证token响应
type OAuthCode2SessionReponse struct {
	OpenID     string `json:"openid"`      // 用户的唯一标识
	Sessionkey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
}

// OAuthGetAccessToken 获取 access_token
// 通过code换取网页授权access_token
func (c *Client) OAuthGetAccessToken(code string) (*OAuthAccessTokenReply, error) {
	url := fmt.Sprintf("%s/sns/oauth2/access_token", c.opts.BaseURL)
	value := map[string]string{
		"code":       code,
		"grant_type": "authorization_code",
	}
	if !c.opts.Proxy {
		value["appid"] = c.opts.AppID
		value["secret"] = c.opts.AppSecret
	}
	result, err := Get(url, value)
	if err != nil {
		return nil, err
	}
	reply := new(OAuthAccessTokenReply)
	json.Unmarshal(result, reply)
	return reply, nil
}

// OAuthRefreshToken 刷新access_token
// 由于access_token拥有较短的有效期，当access_token超时后，可以使用refresh_token进行刷新，refresh_token有效期为30天，当refresh_token失效之后，需要用户重新授权。
func (c *Client) OAuthRefreshToken(accessToken string) (*OAuthRefreshTokenReply, error) {
	url := fmt.Sprintf("%s/sns/oauth2/refresh_token", c.opts.BaseURL)
	value := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": accessToken,
	}
	if !c.opts.Proxy {
		value["appid"] = c.opts.AppID
	}
	result, err := Get(url, value)
	if err != nil {
		return nil, err
	}
	reply := new(OAuthRefreshTokenReply)
	json.Unmarshal(result, reply)
	return reply, nil
}

// OAuthGetUserInfo 拉取用户信息
// 如果网页授权作用域为snsapi_userinfo，则此时开发者可以通过access_token和openid拉取用户信息了。
func (c *Client) OAuthGetUserInfo(accessToken, openID string) (*OAuthUserInfoReply, error) {
	url := fmt.Sprintf("%s/sns/userinfo", c.opts.BaseURL)
	value := map[string]string{
		"openid": openID,
		"lang":   lang.ZH_CN,
	}
	if !c.opts.Proxy {
		value["access_token"] = c.opts.Token.GetAccessToken()
	}
	result, err := Get(url, value)
	if err != nil {
		return nil, err
	}
	reply := new(OAuthUserInfoReply)
	json.Unmarshal(result, reply)
	return reply, nil
}

// OAuthCheckAccessToken 检查Token
func (c *Client) OAuthCheckAccessToken(accessToken, openID string) (bool, error) {
	url := fmt.Sprintf("%s/sns/auth", c.opts.BaseURL)
	value := map[string]string{
		"openid": openID,
	}
	if !c.opts.Proxy {
		value["access_token"] = c.opts.Token.GetAccessToken()
	}
	_, err := Get(url, value)
	if err != nil {
		if err.Error() == "ok" {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

// OAuthCode2Session 小程序登录凭证校验
// 通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (c *Client) OAuthCode2Session(code string) (*OAuthCode2SessionReponse, error) {
	url := fmt.Sprintf("%s/sns/jscode2session", c.opts.BaseURL)
	value := map[string]string{
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	if !c.opts.Proxy {
		value["appid"] = c.opts.AppID
		value["secret"] = c.opts.AppSecret
	}
	result, err := Get(url, value)
	if err != nil {
		return nil, err
	}
	resp := new(OAuthCode2SessionReponse)
	json.Unmarshal(result, resp)
	return resp, nil
}
