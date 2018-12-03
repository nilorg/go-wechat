package menu

import (
	simplejson "github.com/bitly/go-simplejson"
	wechat "github.com/nilorg/go-wechat"
	"github.com/nilorg/go-wechat/context"
)

// Menu 菜单
type Menu struct {
	context *context.Context
}

// NewMenu ...
func NewMenu(c *context.Context) *Menu {
	return &Menu{
		context: c,
	}
}

// Create 创建菜单
func (m *Menu) Create(btns []interface{}) error {
	for i := 0; i < len(btns); i++ {
		if !checkButton(btns[i]) { // 检查类型
			return ErrButtonType
		}
	}
	data := map[string]interface{}{
		"button": btns,
	}

	_, err := wechat.PostJSON("https://api.weixin.qq.com/cgi-bin/menu/create?access_token="+m.context.GetAccessToken(), data)
	if err != nil {
		return err
	}
	return nil
}

// GetAll 获取所有菜单
func (m *Menu) GetAll() (*simplejson.Json, error) {
	bytes, err := wechat.Get("https://api.weixin.qq.com/cgi-bin/menu/get?access_token="+m.context.GetAccessToken(), nil)
	if err != nil {
		return nil, err
	}
	return simplejson.NewJson(bytes)
}

// DeleteAll 删除全部菜单
func (m *Menu) DeleteAll() error {
	_, err := wechat.Get("https://api.weixin.qq.com/cgi-bin/menu/delete?access_token="+m.context.GetAccessToken(), nil)
	if err != nil {
		return err
	}
	return nil
}
