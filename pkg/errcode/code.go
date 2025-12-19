package errcode

type ErrorCode int

const (
	Success      ErrorCode = 0
	ParamError   ErrorCode = 1001
	NotFound     ErrorCode = 1002
	DBError      ErrorCode = 1003
	Unauthorized ErrorCode = 1004
	ServerError  ErrorCode = 1005
)

var codeMsg = map[ErrorCode]string{
	Success:      "OK",
	ParamError:   "参数错误",
	NotFound:     "数据不存在",
	DBError:      "数据库错误",
	Unauthorized: "未授权/未登录",
	ServerError:  "服务内部错误",
}

func (e ErrorCode) Msg() string {
	msg, ok := codeMsg[e]
	if ok {
		return msg
	}
	return codeMsg[ServerError]
}
