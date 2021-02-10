package server

import (
	"github.com/gin-gonic/gin"
	//"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/controller"
	controller "github.com/putriapriandi/cobago/controller"

	//"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/infra"
	infra "github.com/putriapriandi/cobago/infra"

	//"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/util"
	util "github.com/putriapriandi/cobago/util"
)

func Start() {
	server := gin.New()
	server.Use(gin.Recovery(), util.Logger())
	db := infra.LoadOracleDB()
	//db := infra.LoadSQLiteDB()
	inDB := &controller.InDB{DB: db}
	Routes(server, inDB)
	server.Run(":1521")
	defer db.Close()
}
