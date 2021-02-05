package infra

import (
	"github.com/indrahadisetiadi/understanding-go-web-development/model"
	"github.com/jinzhu/gorm"
)

func LoadSQLiteDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&model.User{}, &model.Credential{}, &model.Status{}, &model.UserQuestion{}, &model.UserAnswer{})
	return db
}
