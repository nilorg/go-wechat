package custom

// TextRequest 发送文本消息
type TextRequest struct {
	ToUser  string    `json:"touser"`  // 要发送给的用户openID
	MsgType string    `json:"msgtype"` // 消息类型
	Text    *TextData `json:"text"`    // 消息
}

// TextData 文本消息
type TextData struct {
	Content string `json:"content"` // 消息内容
}

// NewTextRequest 创建发送文本消息
func NewTextRequest(toUser, content string) *TextRequest {
	return &TextRequest{
		ToUser:  toUser,
		MsgType: "text",
		Text: &TextData{
			Content: content,
		},
	}
}

// ImageRequest 发送图片消息
type ImageRequest struct {
	ToUser  string     `json:"touser"`  // 要发送给的用户openID
	MsgType string     `json:"msgtype"` // 消息类型
	Image   *ImageData `json:"image"`   // 消息
}

// ImageData 图片消息
type ImageData struct {
	MediaID string `json:"media_id"` // 素材ID
}

// NewImageRequest 创建发送图片消息
func NewImageRequest(toUser, mediaID string) *ImageRequest {
	return &ImageRequest{
		ToUser:  toUser,
		MsgType: "image",
		Image: &ImageData{
			MediaID: mediaID,
		},
	}
}

// NewsRequest 发送图文消息
type NewsRequest struct {
	ToUser  string    `json:"touser"`  // 要发送给的用户openID
	MsgType string    `json:"msgtype"` // 消息类型
	News    *NewsData `json:"news"`    // 消息
}

// NewsData 图片消息
type NewsData struct {
	Articles []*NewsDataArticle `json:"articles"` // 素材ID
}

// NewsDataArticle 图文消息项
type NewsDataArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

// NewNewsRequest 发送图文消息（点击跳转到外链）
// 图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008。
func NewNewsRequest(toUser string, article *NewsDataArticle) *NewsRequest {
	return &NewsRequest{
		ToUser:  toUser,
		MsgType: "news",
		News: &NewsData{
			Articles: []*NewsDataArticle{article},
		},
	}
}
