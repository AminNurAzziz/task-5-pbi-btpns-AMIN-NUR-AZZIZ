package router

import (
	"project/HINTTASK5/controllers"
	"project/HINTTASK5/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupPhotoRoutes(router *gin.Engine) {
	photoGroup := router.Group("/photo").Use(middlewares.AuthMiddleware())
	{
		photoGroup.POST("/create", controllers.CreatePhoto)
		photoGroup.GET("/get", controllers.GetPhotos)

		// Apply CheckPhotoOwnership middleware for update and delete routes
		photoGroup.PUT("/:id", middlewares.CheckPhotoOwnership(), controllers.UpdatePhoto)
		photoGroup.DELETE("/:id", middlewares.CheckPhotoOwnership(), controllers.DeletePhoto)
	}
}
