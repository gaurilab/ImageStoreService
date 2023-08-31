package main

import (
	"github.com/gaurilab/ImageStoreService/api/handlers"
	"github.com/gaurilab/ImageStoreService/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.RenderHomePage)

	r.GET("/create-album", handlers.RenderCreateAlbumPage)
	r.POST("/create-album", handlers.CreateAlbum)

	r.GET("/delete-album/:albumID", handlers.DeleteAlbum)
	r.GET("/albums/:albumID", handlers.GetAlbum)

	r.GET("/upload/:albumID", handlers.RenderUploadImagePage)
	r.POST("/upload/:albumID", handlers.UploadImage)
	r.GET("/delete-image/:albumID/:imageID", handlers.DeleteImage)
	r.GET("/images/:id", handlers.RenderImage)

	r.Run(":9080")
}
