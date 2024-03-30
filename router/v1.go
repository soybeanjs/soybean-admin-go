package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"soybean-admin-go/api"
	"soybean-admin-go/middleware"
)

func Init(r *gin.Engine) {
	// 使用 cookie 存储会话数据
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("captch"))))
	r.Use(middleware.Cors())
	v1 := r.Group("v1")
	v1.POST("/auth/login", api.Auth.Login)
	v1.Use(middleware.Jwt())
	v1.POST("/auth/logout", api.Auth.GetUserInfo)
}
