package qrcode

// TempQrcodeRequest ，临时二维码
type TempQrcodeRequest struct {
	ExpireSeconds int                    `json:"expire_seconds"` // 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天），此字段如果不填，则默认有效期为30秒。
	ActionName    string                 `json:"action_name"`    // 二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值
	ActionInfo    map[string]interface{} `json:"action_info"`    // 二维码详细信息
}

// TempQrcodeReply 临时二维码回复
type TempQrcodeReply struct {
	Ticket        string `json:"ticket"`         // 获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。
	ExpireSeconds int    `json:"expire_seconds"` // 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天）。
	URL           string `json:"url"`            // 二维码图片解析后的地址，开发者可根据该地址自行生成需要的二维码图片
}

// LimitQrcodeRequest 永久二维码
type LimitQrcodeRequest struct {
	ActionName string                 `json:"action_name"` // 二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值
	ActionInfo map[string]interface{} `json:"action_info"` // 二维码详细信息
}

// LimitQrcodeReply 永久二维码回复
type LimitQrcodeReply struct {
	Ticket string `json:"ticket"` // 获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。
	URL    string `json:"url"`    // 二维码图片解析后的地址，开发者可根据该地址自行生成需要的二维码图片
}

// NewTempStrQrcodeRequest 创建一个临时字符串二维码
// 场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
func NewTempStrQrcodeRequest(str string, expireSeconds int) *TempQrcodeRequest {
	if expireSeconds <= 0 {
		expireSeconds = 30
	}
	return &TempQrcodeRequest{
		ExpireSeconds: expireSeconds,
		ActionName:    "QR_STR_SCENE", // 临时的字符串参数值
		ActionInfo: map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": str,
			},
		},
	}
}

// NewTempQrcodeRequest 创建一个临时二维码
// 场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000）
func NewTempQrcodeRequest(num uint, expireSeconds int) *TempQrcodeRequest {
	if expireSeconds <= 0 {
		expireSeconds = 30
	}
	if num == 0 || num > 100000 {
		panic("临时二维码时为32位非0整型，永久二维码时最大值为100000")
	}
	return &TempQrcodeRequest{
		ExpireSeconds: expireSeconds,
		ActionName:    "QR_SCENE", // 临时的整型参数值
		ActionInfo: map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_id": num,
			},
		},
	}
}

// NewStrLimitQrcodeRequest 创建一个永久字符串二维码
// 场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
func NewStrLimitQrcodeRequest(str string) *LimitQrcodeRequest {
	l := len(str)
	if l == 0 || l > 64 {
		panic("长度限制为1到64")
	}
	return &LimitQrcodeRequest{
		ActionName: "QR_LIMIT_STR_SCENE", // 永久的字符串参数值
		ActionInfo: map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": str,
			},
		},
	}
}

// NewLimitQrcodeRequest 创建一个永久二维码
// 场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000）
func NewLimitQrcodeRequest(num uint) *LimitQrcodeRequest {
	if num == 0 || num > 100000 {
		panic("临时二维码时为32位非0整型，永久二维码时最大值为100000")
	}
	return &LimitQrcodeRequest{
		ActionName: "QR_LIMIT_SCENE", // 永久的整型参数值
		ActionInfo: map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_id": num,
			},
		},
	}
}
