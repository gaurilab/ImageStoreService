package routes

import (
    "github.com/gin-gonic/gin"
    "api/handlers"
)

func SetupAlbumRoutes(router *gin.RouterGroup) {
    albumRoutes := router.Group("/albums")
    {
        albumRoutes.POST("/", handlers.CreateAlbum)
        albumRoutes.DELETE("/:albumID", handlers.DeleteAlbum)
        // Define other album routes here
    }
}

