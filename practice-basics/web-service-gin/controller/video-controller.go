package controller

import (
	"hello/web-service-gin/entity"
	service "hello/web-service-gin/video-service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindALL() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller{
		service: service,
	}
}
func (c *controller) FindALL() []entity.Video {
	return c.service.FindAll()
}
func (c *controller) save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
