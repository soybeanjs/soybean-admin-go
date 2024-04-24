package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"soybean-admin-go/config"
	"soybean-admin-go/db"
	"soybean-admin-go/router"
	"time"
)

func main() {
	var Loc, _ = time.LoadLocation("Asia/Shanghai")
	time.Local = Loc

	app := gin.Default()
	config.Init()
	db.Init()
	router.Init(app)
	err := app.Run(":9528")
	fmt.Printf("%#v", err)
}
