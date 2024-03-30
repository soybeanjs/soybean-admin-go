package inout

type LoginReq struct {
	Username string `form:"userName" binding:"required"`
	Password string `form:"password" binding:"required"`
}
