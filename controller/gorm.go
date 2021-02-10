package controller

import (
	"github.com/jinzhu/gorm"
	//"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/infra"
	infra "github.com/putriapriandi/cobago/infra"
)

type InDB struct {
	DB *gorm.DB
}

func ConnectDB() *InDB {
	//db := infra.LoadPostgreSQLDB()
	db := infra.LoadOracleDB()
	//db := infra.LoadSQLiteDB()
	inDB := &InDB{DB: db}
	return inDB
}
