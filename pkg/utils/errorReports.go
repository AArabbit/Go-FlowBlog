package utils

import "flow-blog/internal/global"

// RecordError 记录错误日志，并panic
func RecordError(msg string, err error) {
	errInfo := msg + err.Error()
	global.Log.Error(errInfo)
	panic(errInfo)
}
