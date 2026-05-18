package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nazala-mafa/go-crud/config"
	"github.com/nazala-mafa/go-crud/dto"
	"github.com/nazala-mafa/go-crud/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		DB: db,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var payload dto.LoginRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var user models.User

	if err := h.DB.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid credentials",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(payload.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid credentials",
		})
		return
	}

	exp := time.Hour * 24

	if payload.Remember {
		exp = time.Hour * 24 * 30
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":   user.ID,
			"email": user.Email,
			"exp":   time.Now().Add(exp).Unix(),
		},
	)

	tokenString, err := token.SignedString(
		[]byte(config.Jwt.Secret),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed generate token",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": tokenString,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	token := strings.TrimPrefix(
		authHeader,
		"Bearer ",
	)

	err := h.DB.Create(&models.RevokedToken{
		Token: token,
	}).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed logout",
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *AuthHandler) Me(c *gin.Context) {
	c.JSON(200, gin.H{
		"user_id": c.MustGet("user_id"),
		"email":   c.MustGet("email"),
	})
}
