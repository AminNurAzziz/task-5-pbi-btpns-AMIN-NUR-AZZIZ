package helpers

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"project/HINTTASK5/config"
)


func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	return token.SignedString([]byte(config.GetJWTSecret()))
}

func GetUserIDFromToken(c *gin.Context) (uint, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return 0, fmt.Errorf("claims not found in context")
	}

	tokenClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("unable to extract token claims")
	}

	userID, ok := tokenClaims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id not found in token claims")
	}

	return uint(userID), nil
}
