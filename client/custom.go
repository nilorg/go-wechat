package client

import "fmt"

// CustomTextRequest 发送文本消息
type CustomTextRequest struct {
	ToUser  string          `json:"touser"`  // 要发送给的用户openID
	MsgType string          `json:"msgtype"` // 消息类型
	Text    *CustomTextData `json:"text"`    // 消息
}

// CustomTextData 文本消息
type CustomTextData struct {
	Content string `json:"content"` // 消息内容
}

// NewCustomTextRequest 创建发送文本消息
func NewCustomTextRequest(toUser, content string) *CustomTextRequest {
	return &CustomTextRequest{
		ToUser:  toUser,
		MsgType: "text",
		Text: &CustomTextData{
			Content: content,
		},
	}
}

// CustomImageRequest 发送图片消息
type CustomImageRequest struct {
	ToUser  string           `json:"touser"`  // 要发送给的用户openID
	MsgType string           `json:"msgtype"` // 消息类型
	Image   *CustomImageData `json:"image"`   // 消息
}

// CustomImageData 图片消息
type CustomImageData struct {
	MediaID string `json:"media_id"` // 素材ID
}

// NewCustomImageRequest 创建发送图片消息
func NewCustomImageRequest(toUser, mediaID string) *CustomImageRequest {
	return &CustomImageRequest{
		ToUser:  toUser,
		MsgType: "image",
		Image: &CustomImageData{
			MediaID: mediaID,
		},
	}
}

// CustomNewsRequest 发送图文消息
type CustomNewsRequest struct {
	ToUser  string          `json:"touser"`  // 要发送给的用户openID
	MsgType string          `json:"msgtype"` // 消息类型
	News    *CustomNewsData `json:"news"`    // 消息
}

// CustomNewsData 图片消息
type CustomNewsData struct {
	Articles []*CustomNewsDataArticle `json:"articles"` // 素材ID
}

// CustomNewsDataArticle 图文消息项
type CustomNewsDataArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

// NewCustomNewsRequest 发送图文消息（点击跳转到外链）
// 图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008。
func NewCustomNewsRequest(toUser string, article *CustomNewsDataArticle) *CustomNewsRequest {
	return &CustomNewsRequest{
		ToUser:  toUser,
		MsgType: "news",
		News: &CustomNewsData{
			Articles: []*CustomNewsDataArticle{article},
		},
	}
}

// CustomMenuRequest 菜单请求
type CustomMenuRequest struct {
	ToUser  string             `json:"touser"`  // 要发送给的用户openID
	MsgType string             `json:"msgtype"` // 消息类型
	Menus   *CustomMenuMessage `json:"msgmenu"` // 菜单
}

// CustomMenuMessage 菜单消息
type CustomMenuMessage struct {
	HeadContent string                       `json:"head_content"` // 菜单名称
	List        []*CustomMenuMessageListItem `json:"list"`         // 菜单项
	TailContent string                       `json:"tail_content"` // 结语
}

// CustomMenuMessageListItem 菜单项
type CustomMenuMessageListItem struct {
	ID      string `json:"id"`      // ID
	Content string `json:"content"` // 内容
}

// NewCustomMenuRequest 发送菜单消息
func NewCustomMenuRequest(toUser string, msg *CustomMenuMessage) *CustomMenuRequest {
	return &CustomMenuRequest{
		ToUser:  toUser,
		MsgType: "msgmenu",
		Menus:   msg,
	}
}

// send 发消息
func (c *Client) customSend(req interface{}) error {
	url := fmt.Sprintf("%s/cgi-bin/message/custom/send", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	_, err := PostJSON(c.opts.HttpClient, url, req)
	return err
}

// SendText 发送文本消息
func (c *Client) CustomSendText(req *CustomTextRequest) error {
	return c.customSend(req)
}

// SendImage 发送图片消息
func (c *Client) CustomSendImage(req *CustomImageRequest) error {
	return c.customSend(req)
}

// SendNews 发送图文消息（点击跳转到外链）
// 图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008。
func (c *Client) CustomSendNews(req *CustomNewsRequest) error {
	return c.customSend(req)
}

// SendMenu 发送菜单消息
func (c *Client) CustomSendMenu(req *CustomMenuRequest) error {
	return c.customSend(req)
}
