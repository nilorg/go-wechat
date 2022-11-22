package client

import (
	"encoding/json"
	"errors"
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
)

// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141013

const (
	// MenuButtonTypeClick ..
	MenuButtonTypeClick = "click"
	// MenuButtonTypeView ...
	MenuButtonTypeView = "view"
	// MenuButtonTypeMiniProgram ...
	MenuButtonTypeMiniProgram = "miniprogram"
	// MenuButtonTypeMediaID ...
	MenuButtonTypeMediaID = "media_id"
	// MenuButtonTypeViewLimited ...
	MenuButtonTypeViewLimited = "view_limited"
)

var (
	// ErrMenuSubButtonAppend ...
	ErrMenuSubButtonAppend = errors.New("追加按钮错误，请查查按钮类型")
	// ErrMenuButtonType ...
	ErrMenuButtonType = errors.New("菜单类型错误")
)

// menuCheckButton ...
func menuCheckButton(x interface{}) bool {
	switch x.(type) {
	case *MenuButton, *MenuClickButton, *MenuViewButton, *MenuMiniProgramButton, *MenuMediaIDButton, *MenuViewLimitedButton:
		return true
	}
	return false
}

// MenuSubButton ...子级按钮类型
type MenuSubButton []interface{}

// Append 添加子按钮
func (s *MenuSubButton) Append(x interface{}) error {
	if !menuCheckButton(x) {
		return ErrMenuSubButtonAppend
	}
	*s = append(*s, x)
	return nil
}

// MenuButton 按钮
type MenuButton struct {
	Type      string        `json:"type"`
	Name      string        `json:"name"`
	Key       string        `json:"key"`
	SubButton MenuSubButton `json:"sub_button"`
}

// JSON ...
func (b *MenuButton) JSON() string {
	bytes, _ := json.Marshal(b)
	return string(bytes)
}

// NewMenuButton ...
func NewMenuButton() *MenuButton {
	return &MenuButton{
		SubButton: MenuSubButton{},
	}
}

// MenuClickButton click点击类型
type MenuClickButton struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

// NewClickButton ...
func NewClickButton() *MenuClickButton {
	return &MenuClickButton{
		Type: MenuButtonTypeClick,
	}
}

// MenuViewButton view类型
type MenuViewButton struct {
	Type string `json:"type"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// NewViewButton ...
func NewViewButton() *MenuViewButton {
	return &MenuViewButton{
		Type: MenuButtonTypeView,
	}
}

// MenuMiniProgramButton 小程序类型
type MenuMiniProgramButton struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

// NewMenuMiniProgramButton ...
func NewMenuMiniProgramButton() *MenuMiniProgramButton {
	return &MenuMiniProgramButton{
		Type: MenuButtonTypeMiniProgram,
	}
}

// MenuMediaIDButton 永久素材类型可以是图片、音频、视频、图文消息
type MenuMediaIDButton struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	MediaID string `json:"media_id"`
}

// NewMenuMediaIDButton ...
func NewMenuMediaIDButton() *MenuMediaIDButton {
	return &MenuMediaIDButton{
		Type: MenuButtonTypeMediaID,
	}
}

// MenuViewLimitedButton 小程序类型
type MenuViewLimitedButton struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	MediaID string `json:"media_id"`
}

// NewMenuViewLimitedButton ...
func NewMenuViewLimitedButton() *MenuViewLimitedButton {
	return &MenuViewLimitedButton{
		Type: MenuButtonTypeViewLimited,
	}
}

// MenuCreate 创建菜单
func (c *Client) MenuCreate(btns []interface{}) error {
	for i := 0; i < len(btns); i++ {
		if !menuCheckButton(btns[i]) { // 检查类型
			return ErrMenuButtonType
		}
	}
	data := map[string]interface{}{
		"button": btns,
	}
	url := fmt.Sprintf("%s/cgi-bin/menu/create", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	_, err := PostJSON(c.opts.HttpClient, url, data)
	if err != nil {
		return err
	}
	return nil
}

// MenuGetAll 获取所有菜单
func (c *Client) MenuGetAll() (*simplejson.Json, error) {
	url := fmt.Sprintf("%s/cgi-bin/menu/get", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	bytes, err := Get(c.opts.HttpClient, url, nil)
	if err != nil {
		return nil, err
	}
	return simplejson.NewJson(bytes)
}

// MenuDeleteAll 删除全部菜单
func (c *Client) MenuDeleteAll() error {
	url := fmt.Sprintf("%s/cgi-bin/menu/delete", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	_, err := Get(c.opts.HttpClient, url, nil)
	if err != nil {
		return err
	}
	return nil
}
