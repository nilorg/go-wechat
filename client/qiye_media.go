package client

import (
	"encoding/json"
	"fmt"
	"io"
)

type QiyeMedialType string

const (
	// QiyeMedialTypeImage 图片（image）
	QiyeMedialTypeImage QiyeMedialType = "image"
	// QiyeMedialTypeVoice 语音（voice）
	QiyeMedialTypeVoice QiyeMedialType = "voice"
	// QiyeMedialTypeVideo 视频（video）
	QiyeMedialTypeVideo QiyeMedialType = "video"
	// QiyeMedialTypeFile 普通文件（file）
	QiyeMedialTypeFile QiyeMedialType = "file"
)

// QiyeMedialUploadResponse 上传文件响应
type QiyeMedialUploadResponse struct {
	Type      QiyeMedialType `json:"type"`
	MediaID   string         `json:"media_id"`
	CreatedAt int64          `json:"created_at"`
}

// QiyeMedialUpload 上传临时素材
// 素材上传得到media_id，该media_id仅三天内有效, media_id在同一企业内应用之间可以共享
func (c *QiyeClient) QiyeMedialUpload(filename string, fileType QiyeMedialType, srcFile io.Reader) (*QiyeMedialUploadResponse, error) {
	url := fmt.Sprintf("%s/cgi-bin/media/upload?type=%s", c.opts.BaseURL, fileType)
	if !c.opts.Proxy {
		url += fmt.Sprintf("&access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := Upload(url, filename, nil, srcFile)
	if err != nil {
		return nil, err
	}
	resp := new(QiyeMedialUploadResponse)
	json.Unmarshal(result, resp)
	return resp, nil
}

// QiyeMedialGet 获取临时素材
func (c *QiyeClient) QiyeMedialGet(mediaID string, dis io.Writer) error {
	url := fmt.Sprintf("%s/cgi-bin/media/get?media_id=%s", c.opts.BaseURL, mediaID)
	if !c.opts.Proxy {
		url += fmt.Sprintf("&access_token=%s", c.opts.Token.GetAccessToken())
	}
	_, err := Download(url, dis)
	return err
}

// QiyeMedialGetJsSdk 获取高清语音素材
// 可以使用本接口获取从JSSDK的uploadVoice接口上传的临时语音素材，格式为speex，16K采样率。该音频比上文的临时素材获取接口（格式为amr，8K采样率）更加清晰，适合用作语音识别等对音质要求较高的业务。
func (c *QiyeClient) QiyeMedialGetJsSdk(mediaID string, dis io.Writer) error {
	url := fmt.Sprintf("%s/cgi-bin/media/get/jssdk?media_id=%s", c.opts.BaseURL, mediaID)
	if !c.opts.Proxy {
		url += fmt.Sprintf("&access_token=%s", c.opts.Token.GetAccessToken())
	}
	_, err := Download(url, dis)
	return err
}
