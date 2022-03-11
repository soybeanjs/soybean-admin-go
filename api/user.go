package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/honghuangdc/soybean-admin-go/api/e"
	db "github.com/honghuangdc/soybean-admin-go/db/sqlc"
	"github.com/honghuangdc/soybean-admin-go/util"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum" lable:"用户名"`
	Password string `json:"password" binding:"required,min=6" lable:"密码"`
	FullName string `json:"full_name" binding:"required" lable:"姓名"`
	Email    string `json:"email" binding:"required_without=Phone,omitempty,email" lable:"邮箱"`
	Phone    string `json:"phone" binding:"required_without=Email,omitempty,phone" lable:"手机号码"`
}

type userResponse struct {
	Username         string    `json:"username"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	PasswordChangeAt time.Time `json:"password_change_at"`
	CreatedAt        time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:         user.Username,
		FullName:         user.FullName,
		Email:            user.Email,
		Phone:            user.Phone,
		PasswordChangeAt: user.PasswordChangeAt,
		CreatedAt:        user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	appG := Gin{C: ctx}
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err)
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err)
		return
	}

	// TODO 检查 邮箱和手机号唯一

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
		Phone:          req.Phone,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				appG.Response(http.StatusForbidden, e.ERROR_USER_USERNAME_EXIST, nil)
				return
			}
		}
		appG.Response(http.StatusInternalServerError, e.ERROR, err)
		return
	}

	resp := newUserResponse(user)
	appG.Response(http.StatusOK, e.SUCCESS, resp)
}
