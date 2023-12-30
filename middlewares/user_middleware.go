package middlewares

import (
	"net/http"
	"project/HINTTASK5/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckUserOwnership() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from the URL parameters
		userIDParam := c.Param("id")

		userID, err := strconv.ParseUint(userIDParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			c.Abort()
			return
		}

		// Get user ID from the token
		tokenUserID, err := helpers.GetUserIDFromToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if the user is the owner of the requested user ID
		if uint(userID) != tokenUserID {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to perform this action"})
			c.Abort()
			return
		}


		// Continue to the next middleware or handler
		c.Next()
	}
}
