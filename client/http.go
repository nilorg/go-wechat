package client

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

// Post send post request.
func Post(httpClient *http.Client, url, contentType string, args map[string]string) (result string, err error) {
	if contentType == "" {
		contentType = "application/x-www-form-urlencoded"
	}
	resp, err := httpClient.Post(url, contentType, strings.NewReader(argsEncode(args)))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = NewError(body)
	if err != nil {
		return
	}
	result = string(body)
	return
}

// VideoDescription 视频描述
type VideoDescription struct {
	Title        string `json:"title"`        // 视频素材的标题
	Introduction string `json:"introduction"` // 视频素材的描述
}

// Upload send post request.
func Upload(httpClient *http.Client, uri, filename string, description *VideoDescription, srcFile io.Reader) (result []byte, err error) {
	return upload(httpClient, uri, "media", filename, description, srcFile)
}

// upload send post request.
func upload(httpClient *http.Client, uri, fieldname, filename string, description *VideoDescription, srcFile io.Reader) (result []byte, err error) {
	buf := new(bytes.Buffer)
	// 文件
	writer := multipart.NewWriter(buf)
	formFile, err := writer.CreateFormFile(fieldname, filename)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(formFile, srcFile); err != nil {
		return nil, err
	}
	contentType := writer.FormDataContentType()
	// 附加参数
	if description != nil {
		jsonBytes, _ := json.Marshal(description)
		writer.WriteField("description", string(jsonBytes))
	}
	writer.Close() // 发送之前必须调用Close()以写入结尾行

	resp, err := httpClient.Post(uri, contentType, buf)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = NewError(body)
	if err != nil {
		return
	}
	result = body
	return
}

// Download 下载非视频文件
func Download(httpClient *http.Client, uri string, dis io.Writer) (result []byte, err error) {
	var resp *http.Response
	resp, err = httpClient.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	if strings.Contains(strings.ToLower(ctype), "application/json") {
		result, err = io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		err = NewError(result)
		if err != nil {
			result = nil
		}
	} else {
		io.Copy(dis, resp.Body)
	}
	return
}

// PostJSON send post request.
func PostJSON(httpClient *http.Client, url string, jsonObject interface{}) (result []byte, err error) {
	buf := new(bytes.Buffer)
	hjson := json.NewEncoder(buf)
	hjson.SetEscapeHTML(false)
	err = hjson.Encode(jsonObject)
	if err != nil {
		return
	}
	resp, err := httpClient.Post(url, "application/json", buf)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = NewError(body)
	if err != nil {
		return
	}
	result = body
	return
}

// Get send get request.
func Get(httpClient *http.Client, url string, args map[string]string) (result []byte, err error) {
	if args != nil {
		url += "?" + argsEncode(args)
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = NewError(body)
	if err != nil {
		return
	}
	result = body
	return
}

func argsEncode(params map[string]string) string {
	args := url.Values{}
	for k, v := range params {
		args.Set(k, v)
	}
	return args.Encode()
}
