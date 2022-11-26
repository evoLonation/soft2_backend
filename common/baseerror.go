package common

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type MsgError struct {
	Msg string `json:"msg"`
}

func NewMsgError(msg string) *MsgError {
	return &MsgError{Msg: msg}
}

func InitHttpErrorHandler() {
	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		return http.StatusPreconditionFailed, NewMsgError(err.Error())
	})
}
