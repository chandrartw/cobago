package controller

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/util"

	"github.com/gin-gonic/gin"
	"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/model"
	"golang.org/x/crypto/bcrypt"
)

func (idb *InDB) CreateUser(ctx *gin.Context) {
	var (
		user model.User
	)

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Please Check Your Data")
		ctx.Abort()
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Credential.Password), 14)
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Bad Password")
		ctx.Abort()
		return
	}
	user.Credential.Password = string(password)
	// user.Status.Code = uuid.NewV4().String()
	b := make([]byte, 32) //equals 8 charachters
	rand.Read(b)
	user.Status.Code = hex.EncodeToString(b)

	// dexp, _ := strconv.Atoi(os.Getenv("VERIFY_DAY_EXPIRED"))
	user.Status.VerifyExp = time.Now().Add(time.Hour * 24)
	err = idb.DB.Create(&user).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Create User")
		ctx.Abort()
		return
	}

	err = util.SendMailVerify(user.Email, user.FirstName, user.Status.Code, user.Status.VerifyExp.Format("2006-01-02 15:04:05"))

	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Send Email")
		ctx.Abort()
		return
	}
	util.ResponseSuccess(ctx, http.StatusOK, user)
}

func (idb *InDB) GetUser(ctx *gin.Context) {
	var (
		user       model.User
		credential model.Credential
		status     model.Status
	)
	id := ctx.Param("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Find Id User")
		ctx.Abort()
		return
	}

	err = idb.DB.Where("user_id = ?", id).First(&credential).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Find Id User")
		ctx.Abort()
		return
	}
	user.Credential = credential

	err = idb.DB.Where("user_id = ?", id).First(&status).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Find Id User")
		ctx.Abort()
		return
	}
	user.Status = status

	util.ResponseSuccess(ctx, http.StatusOK, user)
}

func (idb *InDB) GetAllUser(ctx *gin.Context) {

	users := []model.User{}
	_ = idb.DB.Find(&users).Error
	for i, _ := range users {
		idb.DB.Model(users[i]).Related(&users[i].Credential)
	}
	if len(users) <= 0 {
		util.ResponseSuccessCustomMessage(ctx, http.StatusOK, "No Record")
	} else {
		util.ResponseSuccess(ctx, http.StatusOK, users)
	}

}

func (idb *InDB) DeleteUser(ctx *gin.Context) {
	var (
		user model.User
	)
	id := ctx.Param("id")

	err := idb.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Find Id User")
		ctx.Abort()
		return
	}
	err = idb.DB.Unscoped().Delete(&user).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Find Id User")
		ctx.Abort()
		return
	}
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Delete User")
		ctx.Abort()
		return
	}

	util.ResponseSuccessCustomMessage(ctx, http.StatusOK, "Success Deleted")
}

func (idb *InDB) UpdateUser(ctx *gin.Context) {
	id := ctx.Query("id")

	var (
		user    model.User
		newUser model.User
	)

	err := idb.DB.First(&user, id).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "ID Not Found")
		ctx.Abort()
		return
	}

	err = ctx.ShouldBindJSON(&newUser)
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Please Check Your Data")
		ctx.Abort()
		return
	}

	err = idb.DB.Model(&user).Updates(newUser).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Update Failed")
		ctx.Abort()
		return
	}

	util.ResponseSuccess(ctx, http.StatusOK, user)

}

func (idb *InDB) VerifyUser(ctx *gin.Context) {
	code := ctx.Param("code")

	var (
		user   model.User
		status model.Status
	)

	err := idb.DB.Where("code = ?", code).First(&status).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "ID Not Found")
		ctx.Abort()
		return
	}

	err = idb.DB.Where("id = ?", status.UserID).First(&user).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "ID Not Found")
		ctx.Abort()
		return
	}

	status.Verify = 1
	now := time.Now()
	tomorrow := status.VerifyExp
	flag := now.Before(tomorrow)
	if flag == false {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, "Expired, Please Contact your administrator !", "Expired")
		ctx.Abort()
		return
	}
	status.Code = ""
	err = idb.DB.Model(&status).Updates(status).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Update Failed")
		ctx.Abort()
		return
	}

	util.ResponseSuccessCustomMessage(ctx, http.StatusOK, "Success Verify")

}

func (idb *InDB) RequestForgetPassword(ctx *gin.Context) {

	var (
		user   model.User
		status model.Status
	)

	type InputEmail struct {
		Email string `json:"email" validate:"required,email"`
	}

	var inputEmail InputEmail

	err := ctx.ShouldBindJSON(&inputEmail)
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Please Check Your Data")
		ctx.Abort()
		return
	}

	err = idb.DB.Where("email = ?", inputEmail.Email).First(&user).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Email Not Found")
		ctx.Abort()
		return
	}

	err = idb.DB.Where("user_id = ?", user.ID).First(&status).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Email Not Found")
		ctx.Abort()
		return
	}

	b := make([]byte, 32) //equals 8 charachters
	rand.Read(b)
	status.Code = hex.EncodeToString(b)
	status.IsForget = true
	status.VerifyExp = time.Now().Add(time.Hour * 24)
	err = idb.DB.Model(&status).Updates(status).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Update Failed")
		ctx.Abort()
		return
	}

	err = util.SendMailForget(user.Email, user.FirstName, status.Code, user.Status.VerifyExp.Format("2006-01-02 15:04:05"))

	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Send Email")
		ctx.Abort()
		return
	}
	util.ResponseSuccessCustomMessage(ctx, http.StatusOK, "Success Request")

}

func (idb *InDB) VerifyForgetPassword(ctx *gin.Context) {

	code := ctx.Param("code")

	var (
		status model.Status
	)

	err := idb.DB.Where("code = ?", code).First(&status).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "ID Not Found")
		ctx.Abort()
		return
	}

	now := time.Now()
	tomorrow := status.VerifyExp
	flag := now.Before(tomorrow)
	if flag == false {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, "Expired, Please Contact your administrator !", "Expired")
		ctx.Abort()
		return
	}

	util.ResponseSuccess(ctx, http.StatusOK, status.Code)

}

func (idb *InDB) ProsesForgetPassword(ctx *gin.Context) {
	type InputEmail struct {
		Email          string `json:"email" validate:"required,email"`
		Password       string `json:"password" validate:"required,email"`
		VerifyPassword string `json:"verify_password" validate:"required,email"`
	}
}
