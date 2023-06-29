package controllers

import (
	"github.com/NickYxy/GoTypingTest/gin-example/models"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	GetAll(context *gin.Context)

	Update(context *gin.Context)

	Create(context *gin.Context)

	Delete(context *gin.Context)
}

type controller struct {
	videos []models.Video
}

func NewVideoController() VideoController {
	return &controller{}
}

func (c *controller) GetAll(context *gin.Context) {
	context.JSON(200, c.videos)
}

func (c *controller) Update(context *gin.Context) {
	var videoToUpdate models.Video
	if err := context.ShouldBind(&videoToUpdate); err != nil {
		context.String(400, "bad request %v", err)
	}

	for idx, video := range c.videos {
		if video.Id == videoToUpdate.Id {
			c.videos[idx] = videoToUpdate
			context.String(200, "video with id %d has been updated", videoToUpdate.Id)
		}
	}

	context.String(400, "bad request cannot find videos with %d to update", videoToUpdate.Id)
}

func (c *controller) Create(context *gin.Context) {
	panic(any("aa"))
}

func (c *controller) Delete(context *gin.Context) {
	panic(any("aa"))
}
