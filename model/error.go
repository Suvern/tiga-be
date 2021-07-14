package model

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	Err  error
	Code int
	Msg  string
}

func (err HttpError) Error() string {
	return fmt.Sprintf("%d: %s", err.Code, err.Msg)
}

var (
	WrongPasswordError    = HttpError{Code: http.StatusBadRequest, Msg: "用户名不存在或密码错误"}
	UnauthorizedError     = HttpError{Code: http.StatusUnauthorized, Msg: "未登录或权限不足"}
	UnexpectedParamsError = HttpError{Code: http.StatusBadRequest, Msg: "参数格式错误"}
	UnknownServerError    = HttpError{Code: http.StatusInternalServerError, Msg: "服务器内部错误"}
)
