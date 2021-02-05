package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/xkeyideal/captcha/pool"

	"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/auth"
	"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/controller"
	"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/model"
	"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/util"
)

type profileHandler struct {
	rd auth.AuthInterface
	tk auth.TokenInterface
}

func NewProfile(rd auth.AuthInterface, tk auth.TokenInterface) *profileHandler {
	return &profileHandler{rd, tk}
}

var CaptchaPool = pool.NewCaptchaPool(240, 80, 6, 2, 2, 2)
var cacheBuffer *pool.CaptchaBody

func CaptchaHandler(ctx *gin.Context) {

	var (
		result gin.H
	)
	base_url := GenerateCaptcha()
	result = gin.H{
		"image":        base_url,
		"captcha_code": string(cacheBuffer.Val),
	}
	ctx.JSON(http.StatusOK, result)
	ctx.Abort()

}

func GenerateCaptcha() string {
	cacheBuffer = CaptchaPool.GetImage()
	base_url := base64.StdEncoding.EncodeToString(cacheBuffer.Data.Bytes())
	base_url = "data:image/png;base64," + base_url
	return base_url
}

func CaptchaSolver(ctx *gin.Context) {

	var (
		oCaptcha model.Captcha
	)

	err := ctx.ShouldBindJSON(&oCaptcha)

	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Please check your data")
		ctx.Abort()
		return
	}
	flag := CheckCaptchaSolver(oCaptcha.ValueSolution)
	fmt.Printf(oCaptcha.ValueSolution)
	if flag == true {
		var result map[string]string
		result["result"] = "Success solve the captcha"
		util.ResponseSuccess(ctx, http.StatusOK, result)
		ctx.Abort()
		return
	} else {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Failed solve the captcha")
		ctx.Abort()
		return
	}

}

func CheckCaptchaSolver(valueSolution string) bool {
	if cacheBuffer.Val == nil {
		flag := false
		return flag
	}
	flag := string(cacheBuffer.Val) == valueSolution

	return flag
}

func (h *profileHandler) Login(ctx *gin.Context) {
	type CredentialLogin struct {
		Username      string `json:"username" binding:"required"`
		Password      string `json:"password" binding:"required"`
		ValueSolution string `json:"valueSolution" binding:"required"`
	}

	var user model.User
	var inputCrd CredentialLogin
	var crd model.Credential
	var status model.Status

	err := ctx.ShouldBindJSON(&inputCrd)

	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Please check your data !")
		ctx.Abort()
		return
	}

	crd.Password = inputCrd.Password
	crd.Username = inputCrd.Username
	flag := false
	flag = CheckCaptchaSolver(inputCrd.ValueSolution)

	if flag != true {
		util.ResponseError(ctx, http.StatusBadRequest, "Captcha Not Match", "Failed solve the captcha !")
		ctx.Abort()
		return
	}

	inputPass := inputCrd.Password

	fmt.Printf(crd.Username)

	idb := controller.ConnectDB()

	err = idb.DB.Where("username = ?", crd.Username).First(&crd).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnauthorized, err.Error(), "Username Not Found !")
		ctx.Abort()
		return
	}

	err = idb.DB.Where("id = ?", crd.ID).First(&user).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnauthorized, err.Error(), "Username Not Found !")
		ctx.Abort()
		return
	}

	err = idb.DB.Where("user_id = ?", user.ID).First(&status).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnauthorized, err.Error(), "Username Not Found !")
		ctx.Abort()
		return
	}

	if status.Verify == 0 {
		util.ResponseError(ctx, http.StatusUnauthorized, "Not Verify", "Your Account is not verify yet")
		ctx.Abort()
		return
	}

	if status.IsBlocked == true {
		util.ResponseError(ctx, http.StatusUnauthorized, "Blocked", "Your Account is Blocked, Please Contact Your Administrator")
		ctx.Abort()
		return
	}

	err, msg := PasswordSolver(user.UserType, crd.Username, crd.Password, inputPass)

	if err != nil {
		status.CounterBlock = status.CounterBlock + 1
		if status.CounterBlock > 2 {
			status.IsBlocked = true
			status.BlockedAt = time.Now()
			_ = idb.DB.Model(&status).Updates(status).Error
			util.ResponseError(ctx, http.StatusUnauthorized, err.Error(), "Your Account is Blocked, Please Contact Your Administrator")
			ctx.Abort()
			return
		}
		_ = idb.DB.Model(&status).Updates(status).Error
		util.ResponseError(ctx, http.StatusUnauthorized, msg, "Password Not Match !")
		ctx.Abort()
		return
	}

	ts, err := h.tk.CreateToken(crd.Username)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := h.rd.CreateAuth(crd.Username, ts)
	if saveErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, saveErr.Error())
		return
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	util.ResponseSuccess(ctx, http.StatusOK, tokens)
}

func (h *profileHandler) AuthMiddleware(ctx *gin.Context) {

	metadata, err := h.tk.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		// err := h.rd.DeleteTokens(metadata)
		if err.Error() == "Token is expired" {
			fmt.Printf("ini error gara2 token")
			if metadata != nil {
				_ = h.rd.DeleteTokens(metadata)
			}

		}
		result := gin.H{
			"result": "Invalid Credential",
			"debug":  err.Error(),
			"error":  1,
		}

		ctx.JSON(http.StatusUnauthorized, result)
		ctx.Abort()
		return
	}
	username, err := h.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		ctx.Abort()
		return
	}
	fmt.Println(username)

	//you can proceed to save the  to a database
	ctx.Next()
}

func PasswordSolver(types string, username string, password string, inputPass string) (error, string) {
	var err error
	var msg string
	fmt.Printf(username)
	if types == "LDAP" {
		client := util.NewLDAPClient()
		ok, _, err := client.Authenticate(username, inputPass)
		if err != nil {
			msg = "Invalid Credential LDAP"
			return err, msg
		}
		if !ok {
			msg = "Invalid Credential LDAP"
			return err, msg
		}
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(password), []byte(inputPass))
		if err != nil {
			msg = "Password not match"
			return err, msg
		}
	}
	msg = "Success"
	return err, msg

}

func (h *profileHandler) Logout(c *gin.Context) {
	//If metadata is passed and the tokens valid, delete them from the redis store
	metadata, _ := h.tk.ExtractTokenMetadata(c.Request)
	if metadata != nil {
		deleteErr := h.rd.DeleteTokens(metadata)
		if deleteErr != nil {
			c.JSON(http.StatusBadRequest, deleteErr.Error())
			return
		}
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
