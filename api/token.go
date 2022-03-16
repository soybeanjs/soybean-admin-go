package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/honghuangdc/soybean-admin-go/api/e"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	appg := Gin{C: ctx}
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		appg.Response(http.StatusBadRequest, e.InvalidPrarms, err)
		return
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		appg.Response(http.StatusUnauthorized, e.ErrorAuth, err)
		return
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if err != sql.ErrNoRows {
			appg.Response(http.StatusNotFound, e.NotFound, gin.H{})
			return
		}
		appg.Response(http.StatusInternalServerError, e.Error, err)
		return
	}

	if session.IsBlocked {
		err = fmt.Errorf("blocked seesion")
		appg.Response(http.StatusUnauthorized, e.ErrorAuth, err)
		return
	}

	if session.Username != refreshPayload.Username {
		err = fmt.Errorf("incorrect session user")
		appg.Response(http.StatusUnauthorized, e.ErrorAuth, err)
		return
	}

	if session.RefreshToken != req.RefreshToken {
		err = fmt.Errorf("mismatched seesion token")
		appg.Response(http.StatusUnauthorized, e.ErrorAuth, err)
		return
	}

	if time.Now().After(session.ExpiresAt) {
		err = fmt.Errorf("expired session")
		appg.Response(http.StatusUnauthorized, e.ErrorAuthCheckTokenTimtout, err)
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		appg.Response(http.StatusInternalServerError, e.Error, err)
		return
	}

	rsp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}

	appg.Response(http.StatusOK, e.Success, rsp)
}
