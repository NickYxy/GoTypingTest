package main

import (
	"github.com/NickYxy/GoTypingTest/gin-example/controllers"
	"github.com/NickYxy/GoTypingTest/gin-example/middlewares"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := gin.Default()
	server.Use(middlewares.MyAuth())

	server.GET("/ping", func(context *gin.Context) {
		context.String(200, "%s", "pong")
	})

	server.Static("/resources", "./resources")
	server.StaticFile("/video", "./resources/Myth.mp4")

	videoController := controllers.NewVideoController()
	videoGroup := server.Group("/videos")
	videoGroup.Use(middlewares.MyLogger())

	// GET /videos
	videoGroup.GET("/", videoController.GetAll)

	//PUT /videos/123
	videoGroup.PUT("/:id", videoController.Update)

	// POST /videos
	videoGroup.POST("/", videoController.Create)

	//DELETE /videos/123
	videoGroup.DELETE("/:id", videoController.Delete)

	log.Fatalln(server.Run("localhost:8080"))
}
