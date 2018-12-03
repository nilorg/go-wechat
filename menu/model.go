package menu

import (
	"encoding/json"
	"errors"
)

// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141013

const (
	// ButtonTypeClick ..
	ButtonTypeClick = "click"
	// ButtonTypeView ...
	ButtonTypeView = "view"
	// ButtonTypeMiniProgram ...
	ButtonTypeMiniProgram = "miniprogram"
	// ButtonTypeMediaID ...
	ButtonTypeMediaID = "media_id"
	// ButtonTypeViewLimited ...
	ButtonTypeViewLimited = "view_limited"
)

var (
	// ErrSubButtonAppend ...
	ErrSubButtonAppend = errors.New("追加按钮错误，请查查按钮类型")
	// ErrButtonType ...
	ErrButtonType = errors.New("菜单类型错误")
)

// // checkSubButton ...
// func checkSubButton(x interface{}) bool {
// 	switch x.(type) {
// 	case *ClickButton, *ViewButton, *MiniProgramButton, *MediaIDButton, *ViewLimitedButton:
// 		return true
// 	}
// 	return false
// }

// checkButton ...
func checkButton(x interface{}) bool {
	switch x.(type) {
	case *Button, *ClickButton, *ViewButton, *MiniProgramButton, *MediaIDButton, *ViewLimitedButton:
		return true
	}
	return false
}

// SubButton ...子级按钮类型
type SubButton []interface{}

// Append 添加子按钮
func (s *SubButton) Append(x interface{}) error {
	if !checkButton(x) {
		return ErrSubButtonAppend
	}
	*s = append(*s, x)
	return nil
}

// Button 按钮
type Button struct {
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Key       string    `json:"key"`
	SubButton SubButton `json:"sub_button"`
}

// JSON ...
func (b *Button) JSON() string {
	bytes, _ := json.Marshal(b)
	return string(bytes)
}

// NewButton ...
func NewButton() *Button {
	return &Button{
		SubButton: SubButton{},
	}
}

// ClickButton click点击类型
type ClickButton struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

// NewClickButton ...
func NewClickButton() *ClickButton {
	return &ClickButton{
		Type: ButtonTypeClick,
	}
}

// ViewButton view类型
type ViewButton struct {
	Type string `json:"type"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// NewViewButton ...
func NewViewButton() *ViewButton {
	return &ViewButton{
		Type: ButtonTypeView,
	}
}

// MiniProgramButton 小程序类型
type MiniProgramButton struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

// NewMiniProgramButton ...
func NewMiniProgramButton() *MiniProgramButton {
	return &MiniProgramButton{
		Type: ButtonTypeMiniProgram,
	}
}

// MediaIDButton 永久素材类型可以是图片、音频、视频、图文消息
type MediaIDButton struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	MediaID string `json:"media_id"`
}

// NewMediaIDButton ...
func NewMediaIDButton() *MediaIDButton {
	return &MediaIDButton{
		Type: ButtonTypeMediaID,
	}
}

// ViewLimitedButton 小程序类型
type ViewLimitedButton struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	MediaID string `json:"media_id"`
}

// NewViewLimitedButton ...
func NewViewLimitedButton() *ViewLimitedButton {
	return &ViewLimitedButton{
		Type: ButtonTypeViewLimited,
	}
}
