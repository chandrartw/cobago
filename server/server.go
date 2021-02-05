package server

import (
	"github.com/gin-gonic/gin"
	"github.com/indrahadisetiadi/understanding-go-web-development/controller"
	"github.com/indrahadisetiadi/understanding-go-web-development/infra"
	"github.com/indrahadisetiadi/understanding-go-web-development/util"
)

func Start() {
	server := gin.New()
	server.Use(gin.Recovery(), util.Logger())
	db := infra.LoadPostgreSQLDB()
	// db := infra.LoadSQLiteDB()
	inDB := &controller.InDB{DB: db}
	Routes(server, inDB)
	server.Run(":8080")
	defer db.Close()
}
