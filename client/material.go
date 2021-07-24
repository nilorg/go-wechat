package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

const (
	// MaterialTypeImage 图片（image）
	MaterialTypeImage = "image"
	// MaterialTypeVoice 语音（voice）
	MaterialTypeVoice = "voice"
	// MaterialTypeVideo 视频（video）
	MaterialTypeVideo = "video"
	// MaterialTypeThumb 缩略图（thumb）
	MaterialTypeThumb = "thumb"
)

// MaterialNewsRequest 图文素材
type MaterialNewsRequest struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	ShowCoverPic     bool   `json:"show_cover_pic"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
}

// MaterialNewsReply 图文素材回复
type MaterialNewsReply struct {
	MediaID string `json:"media_id"`
}

// MaterialUploadImgReply 上传图片回复
type MaterialUploadImgReply struct {
	URL string `json:"url"`
}

// MaterialUploadFileReply 上传文件回复
type MaterialUploadFileReply struct {
	MediaID string `json:"media_id"`
	URL     string `json:"url"` // 新增的图片素材的图片URL（仅新增图片素材时会返回该字段）
}

// MaterialUploadTempFileReply 上传临时文件回复
type MaterialUploadTempFileReply struct {
	Type      string `json:"type"`       // 媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb，主要用于视频与音乐格式的缩略图）
	MediaID   string `json:"media_id"`   //	媒体文件上传后，获取标识
	CreatedAt int64  `json:"created_at"` // 媒体文件上传时间戳
}

// MaterialGetTempFileReply 获取临时文件回复
type MaterialGetTempFileReply struct {
	VideoURL string `json:"video_url"` // 如果返回的是视频消息素材
}

// AddNews 新增永久图文素材
func (c *Client2) MaterialAddNews(reqs []*MaterialNewsRequest) (*MaterialNewsReply, error) {
	url := fmt.Sprintf("%s/cgi-bin/material/add_news", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := PostJSON(url, map[string]interface{}{
		"articles": reqs,
	})
	if err != nil {
		return nil, err
	}
	reply := new(MaterialNewsReply)
	json.Unmarshal(result, reply)
	return reply, nil
}

// UploadImg 上传图文消息内的图片获取URL
// 本接口所上传的图片不占用公众号的素材库中图片数量的5000个的限制。图片仅支持jpg/png格式，大小必须在1MB以下。
func (c *Client2) MaterialUploadImg(filename string, srcFile io.Reader) (*MaterialUploadImgReply, error) {
	url := fmt.Sprintf("%s/cgi-bin/media/uploadimg", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := Upload(url, filename, nil, srcFile)
	if err != nil {
		return nil, err
	}
	reply := new(MaterialUploadImgReply)
	json.Unmarshal(result, reply)
	return reply, nil
}

// UploadFile 新增其他类型永久素材
// 通过POST表单来调用接口，表单id为media，包含需要上传的素材内容，有filename、filelength、content-type等信息。请注意：图片素材将进入公众平台官网素材管理模块中的默认分组。
func (c *Client2) MaterialUploadFile(filename, fileType string, description *VideoDescription, srcFile io.Reader) (*MaterialUploadFileReply, error) {
	if fileType == MaterialTypeVideo && description == nil {
		return nil, errors.New("请填写视频素材的描述信息")
	}
	url := fmt.Sprintf("%s/cgi-bin/material/add_material?type=%s", c.opts.BaseURL, fileType)
	if !c.opts.Proxy {
		url += fmt.Sprintf("&access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := Upload(url, filename, description, srcFile)
	if err != nil {
		return nil, err
	}
	reply := new(MaterialUploadFileReply)
	json.Unmarshal(result, reply)
	return reply, nil
}

// UploadTempFile 新增临时素材
// 注意点：
// 1、临时素材media_id是可复用的。
// 2、媒体文件在微信后台保存时间为3天，即3天后media_id失效。
// 3、上传临时素材的格式、大小限制与公众平台官网一致。
// 图片（image）: 2M，支持PNG\JPEG\JPG\GIF格式
// 语音（voice）：2M，播放长度不超过60s，支持AMR\MP3格式
// 视频（video）：10MB，支持MP4格式
// 缩略图（thumb）：64KB，支持JPG格式
func (c *Client2) MaterialUploadTempFile(filename, fileType string, srcFile io.Reader) (*MaterialUploadTempFileReply, error) {
	url := fmt.Sprintf("%s/cgi-bin/media/upload?type=%s", c.opts.BaseURL, fileType)
	if !c.opts.Proxy {
		url += fmt.Sprintf("&access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := Upload(url, filename, nil, srcFile)
	if err != nil {
		return nil, err
	}
	reply := new(MaterialUploadTempFileReply)
	json.Unmarshal(result, reply)
	return reply, nil
}

// GetTempFile 获取临时素材
func (c *Client2) MaterialGetTempFile(mediaID string, dis io.Writer) (*MaterialGetTempFileReply, error) {
	url := fmt.Sprintf("%s/cgi-bin/media/upload?media_id=%s", c.opts.BaseURL, mediaID)
	if !c.opts.Proxy {
		url += fmt.Sprintf("&access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := Download(url, dis)
	if err != nil {
		return nil, err
	}
	if result != nil {
		reply := new(MaterialGetTempFileReply)
		json.Unmarshal(result, reply)
		return reply, nil
	}
	return nil, nil
}
