package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"soybean-admin-go/inout"
)

var Auth = &auth{}

type auth struct {
}

func (*auth) Login(c *gin.Context) {
	var params inout.LoginReq
	err := c.Bind(&params)
	if err != nil {
		Resp.Err(c, "20001", err.Error())
		return
	}
	fmt.Printf("%#v", params)

}

func (*auth) GetUserInfo(c *gin.Context) {

}
