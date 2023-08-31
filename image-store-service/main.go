package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gaurilab/image-store-service/api/handlers"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// User Handlers
	r.POST("/users", handlers.CreateUser)
	r.POST("/users/login", handlers.Login)
	r.POST("/users/logout", handlers.Logout)

	// Album Handlers
	r.POST("/albums", handlers.CreateAlbum)
	r.DELETE("/albums/:id", handlers.DeleteAlbum)
	r.GET("/albums/:id", handlers.GetAlbum)
	r.GET("/albums", handlers.GetAllAlbums)

	// Image Handlers
	r.POST("/albums/:id/images", handlers.UploadImage)
	r.GET("/albums/:id/images", handlers.GetAllImages)
	r.GET("/albums/:id/images/:imageID", handlers.GetImage)
	r.DELETE("/albums/:albumID/images/:imageID", handlers.DeleteImage)

	r.Run(":8080")
}

