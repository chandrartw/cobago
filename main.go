package main

import (
	"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/server"
	"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/util"
)

// var (
// 	videoRepository repository.VideoRepository = repository.NewVideoRepository()
// 	videoService    service.VideoService       = service.New(videoRepository)
// 	videoController controller.VideoController = controller.New(videoService)
// )

func main() {
	util.LoadEnv()
	util.SetupLoggerOutput()
	server.Start()
	// server.GET("/videos", func(ctx *gin.Context) {
	// 	ctx.JSON(200, videoController.FindAll())
	// })

	// server.POST("/videos", func(ctx *gin.Context) {
	// 	err := videoController.Save(ctx)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	} else {
	// 		ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
	// 	}

	// })

	// server.PUT("/videos/:id", func(ctx *gin.Context) {
	// 	err := videoController.Update(ctx)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	} else {
	// 		ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
	// 	}

	// })

	// server.DELETE("/videos/:id", func(ctx *gin.Context) {
	// 	err := videoController.Delete(ctx)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	} else {
	// 		ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
	// 	}

	// })

}
