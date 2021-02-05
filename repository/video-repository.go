package repository

import (
	"github.com/indrahadisetiadi/understanding-go-web-development/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type VideoRepository interface {
	Save(video model.Video)
	Update(video model.Video)
	Delete(video model.Video)
	FindAll() []model.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&model.Video{}, &model.User{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(video model.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video model.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video model.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []model.Video {
	var videos []model.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
