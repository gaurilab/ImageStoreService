package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/gaurilab/image-store-service/api/handlers"
)

func SetupImageRoutes(router *gin.RouterGroup) {
    imageRoutes := router.Group("/images")
    {
        imageRoutes.POST("/:albumID", handlers.UploadImage)
        imageRoutes.DELETE("/:albumID/:imageID", handlers.DeleteImage)
        // Define other image routes here
    }
}

