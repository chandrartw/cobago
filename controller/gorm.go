package controller

import (
	"github.com/jinzhu/gorm"
	"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/infra"
)

type InDB struct {
	DB *gorm.DB
}

func ConnectDB() *InDB {
	// infra.LoadPostgreSQLDB()
	db := infra.LoadPostgreSQLDB()
	// db := infra.LoadSQLiteDB()
	inDB := &InDB{DB: db}
	return inDB
}
