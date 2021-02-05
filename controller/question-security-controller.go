package controller

import (
	"net/http"

	"github.com/indrahadisetiadi/understanding-go-web-development/util"

	"github.com/gin-gonic/gin"
	"github.com/indrahadisetiadi/understanding-go-web-development/model"
)

func (idb *InDB) CreateQuestion(ctx *gin.Context) {
	var (
		question model.UserQuestion
	)

	err := ctx.ShouldBindJSON(&question)
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Please Check Your Data")
		ctx.Abort()
		return
	}

	err = idb.DB.Create(&question).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Create Question")
		ctx.Abort()
		return
	}

	util.ResponseSuccess(ctx, http.StatusOK, question)
}
