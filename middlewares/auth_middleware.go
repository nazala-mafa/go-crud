package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nazala-mafa/go-crud/config"
	"github.com/nazala-mafa/go-crud/models"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "missing token",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			c.Abort()
			return
		}

		var revoked models.RevokedToken

		err := db.Where("token = ?", tokenString).First(&revoked).Error

		if err == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token revoked",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (any, error) {
				return []byte(
					cfg.Jwt.Secret,
				), nil
			},
		)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid claims",
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims["sub"])
		c.Set("email", claims["email"])

		c.Next()
	}
}
