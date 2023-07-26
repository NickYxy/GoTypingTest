package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var limitChain = make(chan struct{}, 10)

func timerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		begin := time.Now()
		ctx.Next()
		fmt.Printf("time elapsed %d ms\n", time.Since(begin).Milliseconds())
	}
}

func limiterMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limitChain <- struct{}{}
		ctx.Next()
		<-limitChain
	}
}
func bizHandler(ctx *gin.Context) {
	time.Sleep(time.Millisecond * 100)
	ctx.String(http.StatusOK, "This is for gin middleware show\n")
}

func main() {
	engine := gin.Default()
	engine.Use(timerMiddleware())
	engine.Use(limiterMiddleware())
	engine.GET("/1", bizHandler)
	engine.GET("/2", bizHandler, bizHandler)
	//engine.GET("/1", timerMiddleware(), limiterMiddleware(), bizHandler)
	//engine.GET("/2", timerMiddleware(), limiterMiddleware(), bizHandler, bizHandler)
	engine.Run("localhost:5678")
}
