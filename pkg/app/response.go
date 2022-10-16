package app

import (
	"example.com/my-gin/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, data any) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg: e.GetMsg(errCode),
		Data: data,
	})
	return
}
