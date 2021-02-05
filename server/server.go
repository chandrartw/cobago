package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/controller"
	"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/infra"
	"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/util"
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
