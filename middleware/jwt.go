package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"soybean-admin-go/api"
	"soybean-admin-go/utils"
)

// JWTAuth 中间件，检查token
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			api.Resp.Err(c, "10002", "请求未携带token，无权限访问")
			c.Abort()
			return
		}
		j := utils.NewJWT()
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				api.Resp.Err(c, "10003", "授权已过期")
				c.Abort()
				return
			}
			api.Resp.Err(c, "10004", err.Error())
			c.Abort()
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("uid", claims.UID)
		c.Next()
	}
}
