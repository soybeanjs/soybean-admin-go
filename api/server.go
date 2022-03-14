package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/honghuangdc/soybean-admin-go/db/sqlc"
	"github.com/honghuangdc/soybean-admin-go/token"
	"github.com/honghuangdc/soybean-admin-go/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	if err := InitTrans("zh"); err != nil {
		log.Fatalf("初始化验证翻译器错误: %s", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
