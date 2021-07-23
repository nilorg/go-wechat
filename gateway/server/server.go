package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	crypter "github.com/heroicyang/wechat-crypter"
	"github.com/nilorg/go-wechat/v2/gateway/middleware"
	"github.com/nilorg/go-wechat/v2/gateway/models"
	"github.com/nilorg/sdk/convert"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func HTTP() {
	engine := gin.Default()
	engine.Use(middleware.Header())
	engine.GET("/:appid", checkAppID, AcceptGET)
	engine.POST("/:appid", checkAppID, AcceptPOST)
	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}

func AcceptGET(ctx *gin.Context) {
	var model models.WechatRequestModel
	if err := ctx.ShouldBindWith(&model, binding.Query); err != nil {
		logrus.Errorf("ctx.ShouldBindWith: %v", err)
		ctx.Status(400)
		return
	}
	appID := ctx.Param("appid")
	token := viper.GetString(fmt.Sprintf("apps.%s.token", appID))
	encodingAESKey := viper.GetString(fmt.Sprintf("apps.%s.encoding_aes_key", appID))
	msgCrypter, err := crypter.NewMessageCrypter(token, encodingAESKey, appID)
	if err != nil {
		logrus.Errorf("crypter.NewMessageCrypter: %v", err)
		ctx.Status(400)
		return
	}
	if model.Signature != msgCrypter.GetSignature(convert.ToString(model.Timestamp), convert.ToString(model.Nonce), "") {
		ctx.Status(400)
		return
	}
	ctx.String(http.StatusOK, model.Echostr)
}

// checkAppID 检查AppID
func checkAppID(ctx *gin.Context) {
	appID := ctx.Param("appid")
	if !existAppID(appID) {
		ctx.Status(400)
		ctx.Abort()
		return
	}
	ctx.Next()
}

func existAppID(appID string) bool {
	apps := viper.GetStringMap("apps")
	if v, ok := apps[appID]; ok && v.(string) == appID {
		return true
	}
	return false
}

// 接收
func AcceptPOST(ctx *gin.Context) {
	appID := ctx.Param("appid")
	value := viper.GetStringMapString(appID)
	if value == nil {
		ctx.Status(400)
		return
	}
	encryptionMethod := viper.GetString(fmt.Sprintf("apps.%s.encryption_method", appID))
	if encryptionMethod == "secure" {
		bytes, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			logrus.Errorf("ioutil.ReadAll: %v", err)
			ctx.Status(400)
			return
		}

		token := viper.GetString(fmt.Sprintf("apps.%s.token", appID))
		encodingAESKey := viper.GetString(fmt.Sprintf("apps.%s.encoding_aes_key", appID))
		dataType := viper.GetString(fmt.Sprintf("apps.%s.data_type", appID))
		msgCrypter, err := crypter.NewMessageCrypter(token, encodingAESKey, appID)
		if err != nil {
			logrus.Errorf("crypter.NewMessageCrypter: %v", err)
			ctx.Status(400)
			return
		}
		// 解析报文数据
		var msg *models.AcceptEncryptMessage
		if dataType == "xml" {
			msg, err = models.AcceptEncryptMessageParseForXML(bytes)
			if err != nil {
				logrus.Errorf("models.AcceptEncryptMessageParseForXML: %v", err)
				ctx.Status(400)
				return
			}
		} else if dataType == "json" {
			msg, err = models.AcceptEncryptMessageParseForJSON(bytes)
			if err != nil {
				logrus.Errorf("models.AcceptEncryptMessageParseForJSON: %v", err)
				ctx.Status(400)
				return
			}
		} else {
			logrus.Errorln("数据格式不正确")
			ctx.Status(400)
			return
		}
		// 解析加密报文数据
		var encryptBytes []byte
		var inAppID string
		encryptBytes, inAppID, err = msgCrypter.Decrypt(string(msg.Encrypt))
		if err != nil {
			logrus.Warnln("加密报文解析失败")
			ctx.Status(400)
			return
		}
		if appID != inAppID {
			logrus.Warnln("加密报文AppID不匹配")
			ctx.Status(400)
			return
		}
		// 输出解密后的报文
		logrus.Infoln(string(encryptBytes))

	} else if encryptionMethod == "compatible" {

	} else {
		bytes, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			logrus.Errorf("ioutil.ReadAll: %v", err)
			ctx.Status(400)
			return
		}
		logrus.Infoln(string(bytes))
	}
}
