package middlewares

import (
	"net/http"
	"project/HINTTASK5/database"
	"project/HINTTASK5/helpers"
	"project/HINTTASK5/models"
	"github.com/gin-gonic/gin"
)

func CheckPhotoOwnership() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get photo ID from the URL parameters
		photoID := c.Param("id")

		var photo models.Photo
		if err := database.DB.First(&photo, photoID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
			c.Abort()
			return
		}

		// Get user ID from the token
		userID, err := helpers.GetUserIDFromToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if the user is the owner of the photo
		if photo.UserID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to perform this action"})
			c.Abort()
			return
		}

		// Set photo information in the context
		c.Set("photo", photo)

		// Continue to the next middleware or handler
		c.Next()
	}
}
