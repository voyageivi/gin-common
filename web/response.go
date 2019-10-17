package web

import (
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
func (g *Gin) Succ(data interface{}) {
	g.C.JSON(200, Response{
		Code: 200,
		Data: data,
	})
	return
}
func (g *Gin) Error(msg string) {
	g.C.JSON(500, Response{
		Code: 500,
		Msg:  msg,
	})
	return
}
