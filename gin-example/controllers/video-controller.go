package controllers

import (
	"github.com/NickYxy/GoTypingTest/gin-example/models"
	"github.com/gin-gonic/gin"
	"sync"
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

type generator struct {
	counter int
	mtx     sync.Mutex
}

func (g *generator) getNextId() int {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	g.counter++
	return g.counter
}

var g *generator = &generator{}

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
	video := models.Video{Id: g.getNextId()}
	if err := context.BindJSON(&video); err != nil {
		context.String(400, "bad request %v", err)
	}
	c.videos = append(c.videos, video)
	context.String(200, "success, new video id is %d", video.Id)
}

func (c *controller) Delete(context *gin.Context) {
	var videoToDelete models.Video
	if err := context.ShouldBind(&videoToDelete); err != nil {
		context.String(400, "bad request %v", err)
	}
	for idx, video := range c.videos {
		if video.Id == videoToDelete.Id {
			c.videos = append(c.videos[0:idx], c.videos[idx+1:len(c.videos)]...)
			context.String(200, "success, video with %d has been deleted", videoToDelete.Id)
		}
	}
	context.String(400, "bad request, cannot find video with %d to delete", videoToDelete.Id)
}
