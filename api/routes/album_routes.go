package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/gauriab/ImageStoreService/api/handlers"
)

func SetupAlbumRoutes(router *gin.RouterGroup) {
    albumRoutes := router.Group("/albums")
    {
        albumRoutes.POST("/", handlers.CreateAlbum)
        albumRoutes.DELETE("/:albumID", handlers.DeleteAlbum)
        // Define other album routes here
    }
}

