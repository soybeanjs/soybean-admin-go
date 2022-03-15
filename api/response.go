package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/honghuangdc/soybean-admin-go/api/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Body interface{} `json:"body"`
}

func (g *Gin) Response(httpCode, errCode int, body interface{}) {
	errs, ok := body.(validator.ValidationErrors)
	if ok {
		rsp := make(map[string]interface{})
		for field, err := range errs.Translate(trans) {
			rsp[field[strings.Index(field, ".")+1:]] = err
		}
		g.C.JSON(httpCode, Response{
			Code: errCode,
			Msg:  e.GetMsg(errCode),
			Body: rsp,
		})
	} else {
		g.C.JSON(httpCode, Response{
			Code: errCode,
			Msg:  e.GetMsg(errCode),
			Body: body,
		})
	}
}

func (g *Gin) AbortWithStatusJSON(httpCode, errCode int, body interface{}) {
	g.C.Abort()
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Body: body,
	})
}
