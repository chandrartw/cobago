package model

import "gorm.io/gorm"

type UserQuestion struct {
	gorm.Model
	Question string `json:"question" gorm:"type:varchar(255)`
}

type UserAnswer struct {
	gorm.Model
	UserID     int    `gorm:"primaryKey"`
	QuestionID int    `gorm:"primaryKey"`
	Answer     string `gorm:"type:varchar(255)"`
}
