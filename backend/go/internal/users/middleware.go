package users

import (
	"net/http"
	"strings"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/auth"
	"github.com/gin-gonic/gin"
)



func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		var token string

		cookieToken, err := c.Cookie("token")
		if err != nil && cookieToken != ""{
			token = cookieToken
		}

		if token == ""{
			authHeader := c.GetHeader("Authorization")

			if strings.HasPrefix(authHeader, "Bearer "){
				token = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if token == ""{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "missing authentication token",
			})
			return
		}

		claims, err := auth.ValidateToken(token)
		if err != nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid or expired token",
			})
			return
		}

		c.Set("user_id",claims.UserID)
		
		c.Next()
	}
}