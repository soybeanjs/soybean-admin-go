package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Resp = &rps{}

type rps struct {
	Code string      `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func (rps) Succ(c *gin.Context, data interface{}) {
	resp := rps{
		Code: "0000",
		Msg:  "请求成功",
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}

func (rps) Err(c *gin.Context, ErrCode, messge string) {
	resp := rps{
		Code: ErrCode,
		Msg:  messge,
	}
	c.JSON(http.StatusOK, resp)
}
