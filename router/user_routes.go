package router

import (
    "project/HINTTASK5/controllers"
    "project/HINTTASK5/middlewares"
    "github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
    userGroup := router.Group("/users")
    {
        // Public routes (no authentication required)
        userGroup.POST("/register", controllers.RegisterUser)
        userGroup.POST("/login", controllers.LoginUser)

        // Authenticated routes (requires authentication)
        userGroupAuthenticated := userGroup.Use(middlewares.AuthMiddleware(), middlewares.CheckUserOwnership())
        {
            userGroupAuthenticated.PUT("/:id", controllers.UpdateUser)
            userGroupAuthenticated.DELETE("/:id", controllers.DeleteUser)
        }
    }
}
