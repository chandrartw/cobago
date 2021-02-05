package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/indrahadisetiadi/understanding-go-web-development/model"
	"github.com/indrahadisetiadi/understanding-go-web-development/service"

	"gopkg.in/go-playground/validator.v9"
)

type VideoController interface {
	FindAll() []model.Video
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []model.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video model.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	c.service.Save(video)
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var video model.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	c.service.Update(video)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var video model.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	c.service.Delete(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
