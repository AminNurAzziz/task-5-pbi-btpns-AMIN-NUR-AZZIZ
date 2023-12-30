package helpers

import (
	"strings"
	"github.com/gin-gonic/gin"
)

func ExtractBearerToken(c *gin.Context) string {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return ""
	}

	return strings.Replace(tokenString, "Bearer ", "", 1)
}
