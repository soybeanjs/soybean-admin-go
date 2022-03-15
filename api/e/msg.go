package e

var MsgFlags = map[int]string{
	Success:           "ok",
	Error:             "fail",
	InvalidPrarms:     "请求参数错误",
	NotFound:          "资源未找到",
	Unauthorized:      "认证失败",
	ErrorUsernameExit: "用户名已存在",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimtout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
