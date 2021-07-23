package models

import (
	"encoding/json"
	"encoding/xml"

	"github.com/nilorg/go-wechat/v2/pkg/cdata"
)

// ResponseEncryptMessage 响应加密消息
type ResponseEncryptMessage struct {
	XMLName      xml.Name    `xml:"xml"`
	Encrypt      cdata.CDATA `xml:"Encrypt" json:"Encrypt"`
	MsgSignature cdata.CDATA `xml:"MsgSignature" json:"MsgSignature"`
	TimeStamp    cdata.CDATA `xml:"TimeStamp" json:"TimeStamp"`
	Nonce        cdata.CDATA `xml:"Nonce" json:"Nonce"`
}

// ResponseEncryptMessageParseForXML ...
func ResponseEncryptMessageParseForXML(xmlValue []byte) (msg *ResponseEncryptMessage, err error) {
	msg = new(ResponseEncryptMessage)
	if err = xml.Unmarshal(xmlValue, msg); err != nil {
		msg = nil
		return
	}
	return
}

// ResponseEncryptMessageParseForJSON ...
func ResponseEncryptMessageParseForJSON(jsonValue []byte) (msg *ResponseEncryptMessage, err error) {
	msg = new(ResponseEncryptMessage)
	if err = json.Unmarshal(jsonValue, msg); err != nil {
		msg = nil
		return
	}
	return
}

// AcceptEncryptMessage 接收加密消息
type AcceptEncryptMessage struct {
	XMLName    xml.Name    `xml:"xml"`
	Encrypt    cdata.CDATA `xml:"Encrypt" json:"Encrypt"`
	ToUserName cdata.CDATA `xml:"ToUserName" json:"ToUserName"`
}

// AcceptEncryptMessageParseForXML ...
func AcceptEncryptMessageParseForXML(xmlValue []byte) (msg *AcceptEncryptMessage, err error) {
	msg = new(AcceptEncryptMessage)
	if err = xml.Unmarshal(xmlValue, msg); err != nil {
		msg = nil
		return
	}
	return
}

// AcceptEncryptMessageParseForJSON ...
func AcceptEncryptMessageParseForJSON(jsonValue []byte) (msg *AcceptEncryptMessage, err error) {
	msg = new(AcceptEncryptMessage)
	if err = json.Unmarshal(jsonValue, msg); err != nil {
		msg = nil
		return
	}
	return
}
