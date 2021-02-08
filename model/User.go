package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName    string       `json:"firstname" binding:"required" gorm:"type:varchar(20)"`
	LastName     string       `json:"lastname" binding:"required" gorm:"type:varchar(20)"`
	Email        string       `json:"email" validate:"required,email" gorm:"type:varchar(150);NOT NULL;UNIQUE;index"`
	UserType     string       `json:"type" validate:"required" gorm:"type:varchar(10)"`
	Credential   Credential   `json:"credential" binding:"required" gorm:"FOREIGNKEY:UserID;ASSOCIATION_FOREIGNKEY:ID;constraint:OnDelete:CASCADE;"`
	CredentialID uint64       `json:"-"`
	Status       Status       `json:"status" binding:"required" gorm:"FOREIGNKEY:UserID;ASSOCIATION_FOREIGNKEY:ID;constraint:OnDelete:CASCADE;"`
	StatusID     uint64       `json:"-"`
	UserAnswer   []UserAnswer `json:"user_answer" binding:"required" gorm:"FOREIGNKEY:UserID;constraint:OnDelete:CASCADE;"`
}

type Credential struct {
	gorm.Model
	UserID   uint64 `json:"user_id"`
	Username string `json:"username" binding:"required" gorm:"type:varchar(20);NOT NULL;UNIQUE;index"`
	Password string `json:"password" binding:"required" gorm:"type:varchar(255)"`
}

type Status struct {
	gorm.Model
	UserID        uint64    `json:"user_id"`
	IsTGPASS      bool      `json:"is_tgpass" gorm:"default:false"`
	IsBlocked     bool      `json:"is_blocked" gorm:"default:false"`
	StartValid    time.Time `json:"start_valid" gorm:"default:CURRENT_TIMESTAMP"`
	EndValid      time.Time `json:"end_valid" binding:"required"`
	CounterBlock  int8      `json:"counter_block" gorm:"default:0"`
	CounterForget int8      `json:"counter_forget" gorm:"default:0"`
	BlockedAt     time.Time `json:"blocked_at" gorm:"default:null"`
	Verify        int8      `json:"verify" gorm:"default:0"`
	VerifyExp     time.Time `json:"verify_expired" gorm:"default:null"`
	IsForget      bool      `json:"is_forget" gorm:"default:false"`
	Code          string    `json:"code" gorm:"default:null"`
}

func (user *User) AfterDelete(tx *gorm.DB) error {
	var crd Credential
	var status Status
	var answer UserAnswer
	err := tx.Unscoped().Where("user_id = ?", user.ID).Delete(&crd).Error
	err = tx.Unscoped().Where("user_id = ?", user.ID).Delete(&status).Error
	err = tx.Unscoped().Where("user_id = ?", user.ID).Delete(&answer).Error
	// fmt.Printf("diekse")
	return err
}
