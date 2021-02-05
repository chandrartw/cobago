package service

import (
	"github.com/indrahadisetiadi/understanding-go-web-development/model"
	"github.com/indrahadisetiadi/understanding-go-web-development/repository"
)

type VideoService interface {
	Save(model.Video) error
	Update(model.Video) error
	Delete(model.Video) error
	FindAll() []model.Video
}

type videoService struct {
	repository repository.VideoRepository
}

func New(videoRepository repository.VideoRepository) VideoService {
	return &videoService{
		repository: videoRepository,
	}
}

func (service *videoService) Save(video model.Video) error {
	service.repository.Save(video)
	return nil
}

func (service *videoService) Update(video model.Video) error {
	service.repository.Update(video)
	return nil
}

func (service *videoService) Delete(video model.Video) error {
	service.repository.Delete(video)
	return nil
}

func (service *videoService) FindAll() []model.Video {
	return service.repository.FindAll()
}
