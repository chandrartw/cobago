package infra

import (
	"github.com/indrahadisetiadi/understanding-go-web-development/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func LoadPostgreSQLDB() *gorm.DB {

	db, err := gorm.Open("postgres", "host=10.1.8.115 port=5435 user=postgres dbname=nprm_db password=Sigmaess#2020 sslmode=disable")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&model.User{}, &model.Credential{}, &model.Status{}, &model.UserQuestion{}, &model.UserAnswer{})
	return db
}
