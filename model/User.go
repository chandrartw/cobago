package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// type User struct {
// 	gorm.Model
// 	FirstName    string       `json:"firstname" binding:"required" gorm:"type:varchar(20)"`
// 	LastName     string       `json:"lastname" binding:"required" gorm:"type:varchar(20)"`
// 	Email        string       `json:"email" validate:"required,email" gorm:"type:varchar(150);NOT NULL;UNIQUE;index"`
// 	UserType     string       `json:"type" validate:"required" gorm:"type:varchar(10)"`
// 	Credential   Credential   `json:"credential" binding:"required" gorm:"FOREIGNKEY:UserID;ASSOCIATION_FOREIGNKEY:ID;constraint:OnDelete:CASCADE;"`
// 	CredentialID uint64       `json:"-"`
// 	Status       Status       `json:"status" binding:"required" gorm:"FOREIGNKEY:UserID;ASSOCIATION_FOREIGNKEY:ID;constraint:OnDelete:CASCADE;"`
// 	StatusID     uint64       `json:"-"`
// 	UserAnswer   []UserAnswer `json:"user_answer" binding:"required" gorm:"FOREIGNKEY:UserID;constraint:OnDelete:CASCADE;"`
// }

// type Credential struct {
// 	gorm.Model
// 	UserID   uint64 `json:"user_id"`
// 	Username string `json:"username" binding:"required" gorm:"type:varchar(20);NOT NULL;UNIQUE;index"`
// 	Password string `json:"password" binding:"required" gorm:"type:varchar(255)"`
// }

// type Status struct {
// 	gorm.Model
// 	UserID        uint64    `json:"user_id"`
// 	IsTGPASS      bool      `json:"is_tgpass" gorm:"default:false"`
// 	IsBlocked     bool      `json:"is_blocked" gorm:"default:false"`
// 	StartValid    time.Time `json:"start_valid" gorm:"default:CURRENT_TIMESTAMP"`
// 	EndValid      time.Time `json:"end_valid" binding:"required"`
// 	CounterBlock  int8      `json:"counter_block" gorm:"default:0"`
// 	CounterForget int8      `json:"counter_forget" gorm:"default:0"`
// 	BlockedAt     time.Time `json:"blocked_at" gorm:"default:null"`
// 	Verify        int8      `json:"verify" gorm:"default:0"`
// 	VerifyExp     time.Time `json:"verify_expired" gorm:"default:null"`
// 	IsForget      bool      `json:"is_forget" gorm:"default:false"`
// 	Code          string    `json:"code" gorm:"default:null"`
// }

// func (user *User) AfterDelete(tx *gorm.DB) error {
// 	var crd Credential
// 	var status Status
// 	var answer UserAnswer
// 	err := tx.Unscoped().Where("user_id = ?", user.ID).Delete(&crd).Error
// 	err = tx.Unscoped().Where("user_id = ?", user.ID).Delete(&status).Error
// 	err = tx.Unscoped().Where("user_id = ?", user.ID).Delete(&answer).Error
// 	// fmt.Printf("diekse")
// 	return err
//}

type Customer struct {
	gorm.Model
	CaNo        		string    				`json:"ca_no" binding:"required" gorm:"type:varchar(20)"`
	BpNo				string					`json:"bp_no"`// foreign key
	CreatedDate     	string      			`json:"created_date" gorm:"type:varchar(20)"`
	CreatedBy     		string      			`json:"created_by" gorm:"type:varchar(20)"`
	ValidTo		    	time.Time 				`json:"valid_to" gorm:"default:CURRENT_TIMESTAMP"`
	CaType				string					`json:"ca_type" binding:"required" gorm:"type:varchar(20)" `
	CaName				string					`json:"ca_name" binding:"required" gorm:"type:varchar(20)" `
	Cca					string					`json:"cca" binding:"required" gorm:"type:varchar(20)" `
	BusinessArea		string					`json:"business_area" binding:"required" gorm:"type:varchar(20)" `
	BpRelation			string					`json:"bp_relation" binding:"required" gorm:"type:varchar(20)" `
	TradingPartner		string					`json:"trading_partner" binding:"required" gorm:"type:varchar(20)" `
	Currency			string					`json:"currency" binding:"required" gorm:"type:varchar(20)" `
	AuthGroup			string					`json:"auth_group" binding:"required" gorm:"type:varchar(20)" `
	RefNo				string					`json:"refno" binding:"required" gorm:"type:varchar(20)" `
	PayTerms			string					`json:"pay_terms" binding:"required" gorm:"type:varchar(20)" `
	ToleranceGrp		string					`json:"tolerance_grp" binding:"required" gorm:"type:varchar(20)" `
	ClearingCat			string					`json:"clearing_cat" binding:"required" gorm:"type:varchar(20)" `
	AccDeterm			string					`json:"acc_Determ" binding:"required" gorm:"type:varchar(20)" `
	PlanGrp				string					`json:"plan_grp" binding:"required" gorm:"type:varchar(20)" `
	InterestKey			string					`json:"interest_key" binding:"required" gorm:"type:varchar(20)" `
	KodeSentral			string					`json:"kode_sentral" binding:"required" gorm:"type:varchar(20)" `
	KodeCatel			string					`json:"kode_catel" binding:"required" gorm:"type:varchar(20)" `
	PayerBp				string					`json:"payer_bp" binding:"required" gorm:"type:varchar(20)" `
	PayerCA				string					`json:"payer_ca" binding:"required" gorm:"type:varchar(20)" `
	IsolirFree     		bool      				`json:"isolir_free" gorm:"default:false"`
	VatFree      		bool      				`json:"vat_free" gorm:"default:false"`
	StamDutyFree     	bool      				`json:"stamp_duty_free" gorm:"default:false"`
	Ppn5pct      		bool      				`json:"ppn_5Pct" gorm:"default:false"`
	DunningGrp		    bool      				`json:"dunning_grp" gorm:"default:false"`

}
