package custom

// TextRequest 发送文本消息
type TextRequest struct {
	ToUser  string    `json:"touser"`  // 要发送给的用户openID
	MsgType string    `json:"msgtype"` // 消息类型
	Text    *TextData `json:"text"`    // 消息类型
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
	Image   *ImageData `json:"image"`   // 消息类型
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
