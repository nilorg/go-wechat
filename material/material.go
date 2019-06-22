package material

import (
	"encoding/json"
	"errors"
	"io"

	wechat "github.com/nilorg/go-wechat"
)

// Material 素材
type Material struct {
	client *wechat.Client
}

// NewMaterial ...
func NewMaterial(c *wechat.Client) *Material {
	return &Material{
		client: c,
	}
}

// AddNews 新增永久图文素材
func (m *Material) AddNews(reqs []*NewsRequest) (*NewsReply, error) {
	result, err := wechat.PostJSON("https://api.weixin.qq.com/cgi-bin/material/add_news?access_token="+m.client.GetAccessToken(), map[string]interface{}{
		"articles": reqs,
	})
	if err != nil {
		return nil, err
	}
	reply := new(NewsReply)
	json.Unmarshal(result, reply)
	return reply, nil
}

// UploadImg 上传图文消息内的图片获取URL
// 本接口所上传的图片不占用公众号的素材库中图片数量的5000个的限制。图片仅支持jpg/png格式，大小必须在1MB以下。
func (m *Material) UploadImg(filename string, srcFile io.Reader) (*UploadImgReply, error) {
	result, err := wechat.Upload("https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token="+m.client.GetAccessToken(), filename, nil, srcFile)
	if err != nil {
		return nil, err
	}
	reply := new(UploadImgReply)
	json.Unmarshal(result, reply)
	return reply, nil
}

// UploadFile 新增其他类型永久素材
// 通过POST表单来调用接口，表单id为media，包含需要上传的素材内容，有filename、filelength、content-type等信息。请注意：图片素材将进入公众平台官网素材管理模块中的默认分组。
func (m *Material) UploadFile(filename, fileType string, description *wechat.VideoDescription, srcFile io.Reader) (*UploadFileReply, error) {
	if fileType == TypeVideo && description == nil {
		return nil, errors.New("请填写视频素材的描述信息")
	}

	result, err := wechat.Upload("https://api.weixin.qq.com/cgi-bin/material/add_material?access_token="+m.client.GetAccessToken()+"&type="+fileType, filename, description, srcFile)
	if err != nil {
		return nil, err
	}
	reply := new(UploadFileReply)
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
func (m *Material) UploadTempFile(filename, fileType string, srcFile io.Reader) (*UploadTempFileReply, error) {
	result, err := wechat.Upload("https://api.weixin.qq.com/cgi-bin/media/upload?access_token="+m.client.GetAccessToken()+"&type="+fileType, filename, nil, srcFile)
	if err != nil {
		return nil, err
	}
	reply := new(UploadTempFileReply)
	json.Unmarshal(result, reply)
	return reply, nil
}
