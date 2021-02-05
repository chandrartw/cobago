package controller

import (
	"github.com/indrahadisetiadi/understanding-go-web-development/infra"
	"github.com/jinzhu/gorm"
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
