package main

import (
	"github.com/LuigiJoseph/go-api.git/controller"
	"github.com/LuigiJoseph/go-api.git/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(400, VideoController.Save(ctx))
	})
	server.Run(":8080")
}
