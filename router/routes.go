package router

import (
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Setup routes for users
    SetupUserRoutes(router)

    // Setup routes for photos
    SetupPhotoRoutes(router)

    return router
}
