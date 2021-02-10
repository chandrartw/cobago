package infra

import (
	"github.com/jinzhu/gorm"
	goracle "gopkg.in/goracle.v2"
	//"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/model"
	model "github.com/putriapriandi/cobago/model"
)

func LoadOracleDB() *gorm.DB {
	username := "STGTREMS"
	password := "Stg_Trems@123"
	host := "10.60.185.186"
	database := "TIBSRET20"

	db, err := gorm.Open("goracle", username+"/"+password+"@"+host+"/"+database)
	if err != nil {
		panic(err.Error())
	}
	//db.AutoMigrate(&model.User{}, &model.Credential{}, &model.Status{}, &model.UserQuestion{}, &model.UserAnswer{})
	db.AutoMigrate(&model.Customer{})

	return db
}
