package api

import (
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/honghuangdc/soybean-admin-go/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	if err := InitTrans("zh"); err != nil {
		log.Fatalf("初始化验证翻译器错误: %s", err)
	}
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
