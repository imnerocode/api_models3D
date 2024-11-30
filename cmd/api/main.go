package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/imnerocode/apis/api_models3D/internal/database"
	"github.com/imnerocode/apis/api_models3D/internal/handlers"
	"github.com/imnerocode/apis/api_models3D/internal/repositories"
	"github.com/imnerocode/apis/api_models3D/internal/services"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	db := database.NewConnectionDB()

	client, err := db.GetConnectionDB()
	if err != nil {
		panic(err)
	}
	dbCtx := db.GetDatabase()
	bucket, err := db.GetBucketFs(dbCtx)
	if err != nil {
		panic(err)
	}

	model3dRepository := repositories.NewModel3DRepository(client)
	model3dService := services.NewModel3DService(model3dRepository, bucket)
	model3dHandler := handlers.NewHandlerModel3D(model3dService)

	r.POST("api/model3d/post", model3dHandler.PostModel3D)
	r.POST("api/model3d/upload", model3dHandler.UploadFile)
	r.GET("api/model3d/get", model3dHandler.GetModel3D)
	r.GET("api/model3d/download", model3dHandler.DownloadFile)
	r.DELETE("api/model3d/delete", model3dHandler.DeleteModel3D)

	r.Run(":8080")

}
