package utils

/*
自定义Error
*/

const (
	StatusBadRequest = 100 // 请求错误
	StatusErrorParse = 110 // 内容解析错误
	StatusOK         = 200 // 请求成功
	StatusRetryLimit = 300 // 重试请求次数已达到
)

type customErr struct {
	code int
	msg  string
}

func (e customErr) Error() string {
	return e.msg
}

func ErrorWithCodeMsg(code int, msg string) customErr {
	return customErr{code: code, msg: msg}
}
