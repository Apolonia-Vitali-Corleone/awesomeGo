// Package e pkg/e/code.go
package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	// 用户错误
	ErrorUserExist      = 10001
	ErrorUserNotExist   = 10002
	ErrorPasswordWrong  = 10003
	ErrorTokenGenerate  = 10004
	ErrorTokenInvalid   = 10005
	ErrorTokenExpired   = 10006
	ErrorPermissionDenied = 10007

	// 数据库错误
	ErrorDatabase      = 20001
	ErrorRecordNotFound = 20002
	ErrorRecordExist   = 20003

	// 业务错误
	ErrorBusiness     = 30001
	ErrorUpload       = 30002
	ErrorValidation   = 30003

	// 授权相关错误码
	CodeUnauthorized     = 401  // 未授权
	CodeForbidden        = 403  // 禁止访问
	CodeTokenExpired     = 402  // Token过期
	CodeTokenInvalid     = 403  // Token无效
)

func init() {
	// 补充错误码对应的消息
	MsgFlags[CodeUnauthorized] = "未授权访问"
	MsgFlags[CodeForbidden] = "禁止访问"
	MsgFlags[CodeTokenExpired] = "Token已过期"
	MsgFlags[CodeTokenInvalid] = "Token无效"
}

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",

	ErrorUserExist:      "用户已存在",
	ErrorUserNotExist:   "用户不存在",
	ErrorPasswordWrong:  "密码错误",
	ErrorTokenGenerate:  "令牌生成失败",
	ErrorTokenInvalid:   "无效的令牌",
	ErrorTokenExpired:   "令牌已过期",
	ErrorPermissionDenied: "权限不足",

	ErrorDatabase:      "数据库错误",
	ErrorRecordNotFound: "记录不存在",
	ErrorRecordExist:   "记录已存在",

	ErrorBusiness:     "业务错误",
	ErrorUpload:       "上传失败",
	ErrorValidation:   "数据验证失败",
}

// GetMsg 获取错误信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}