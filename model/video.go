package model

import "time"

type Video struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Title       string    `json:"title" binding:"min=2,max=50" gorm:"type:varchar(50)"`
	Description string    `json:"description" binding:"max=100" gorm:"type:varchar(100)"`
	URL         string    `json:"url" binding:"required" gorm:"type:varchar(20);UNIQUE"`
	Author      User      `json:"author" binding:"required" gorm:"foreignkey:UserID"`
	UserID      uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
