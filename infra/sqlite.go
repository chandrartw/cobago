package infra

import (
	"github.com/jinzhu/gorm"
	"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/model"
)

func LoadSQLiteDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&model.User{}, &model.Credential{}, &model.Status{}, &model.UserQuestion{}, &model.UserAnswer{})
	return db
}
