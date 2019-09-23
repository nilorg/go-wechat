package material

// NewsRequest 图文素材
type NewsRequest struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	ShowCoverPic     bool   `json:"show_cover_pic"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
}

// NewsReply 图文素材回复
type NewsReply struct {
	MediaID string `json:"media_id"`
}

// UploadImgReply 上传图片回复
type UploadImgReply struct {
	URL string `json:"url"`
}

// UploadFileReply 上传文件回复
type UploadFileReply struct {
	MediaID string `json:"media_id"`
	URL     string `json:"url"` // 新增的图片素材的图片URL（仅新增图片素材时会返回该字段）
}

// UploadTempFileReply 上传临时文件回复
type UploadTempFileReply struct {
	Type      string `json:"type"`       // 媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb，主要用于视频与音乐格式的缩略图）
	MediaID   string `json:"media_id"`   //	媒体文件上传后，获取标识
	CreatedAt int64  `json:"created_at"` // 媒体文件上传时间戳
}

// GetTempFileReply 获取临时文件回复
type GetTempFileReply struct {
	VideoURL string `json:"video_url"` // 如果返回的是视频消息素材
}
