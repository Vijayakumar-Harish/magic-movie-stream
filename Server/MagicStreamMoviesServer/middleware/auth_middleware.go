package middleware

import (
	"net/http"

	util "github.com/Vijayakumar-Harish/MagicStreamMovies/Server/MagicStreamMoviesServer/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := util.GetAccessToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
			c.Abort()
			return
		}
		if token == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"No token provided"})
			c.Abort()
			return
		}
		claims, err := util.ValidateToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid token"})
			c.Abort()
			return
		}
		c.Set("userId", claims.UserId)
		c.Set("role", claims.Role)
		
		c.Next()
	}
}