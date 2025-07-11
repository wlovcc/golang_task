package main

import (
	"Task_04/api"
	"Task_04/config"
	"net/http"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}

var (
	ErrDB         = &APIError{Code: http.StatusInternalServerError, Message: "数据库连接失败"}
	ErrAuth       = &APIError{Code: http.StatusUnauthorized, Message: "认证失败，请先登录"}
	ErrNotFound   = &APIError{Code: http.StatusNotFound, Message: "资源不存在"}
	ErrBadRequest = &APIError{Code: http.StatusBadRequest, Message: "请求参数错误"}
)

func main() {
	config.DB = config.InitDb("root", "123456", "127.0.0.1", 3306, "godb")
	api.InitGin()
}
