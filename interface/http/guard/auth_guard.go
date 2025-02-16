package guard

import (
	"mini-wallet-exercise/entities"
	"mini-wallet-exercise/interface/http/exception"
	dbconn "mini-wallet-exercise/pkg/db"
	"strings"

	"github.com/gin-gonic/gin"
)

// Guard to get Authorization: Token from request header
func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			panic(*exception.UnauthorizedException("Authorization header required"))
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Token" {
			panic(*exception.UnauthorizedException("Invalid token format"))
		}

		var customer entities.CustomerEntity
		result := dbconn.DB.Where("token = ?", tokenParts[1]).First(&customer)
		if result.Error != nil {
			panic(*exception.UnauthorizedException("Invalid token"))
		}

		c.Set("customer", customer)
		c.Next()
	}

}
