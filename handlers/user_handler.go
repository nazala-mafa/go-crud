package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nazala-mafa/go-crud/config"
	"github.com/nazala-mafa/go-crud/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

type UserResponse struct {
	models.User
	Link string `json:"link"`
}

func (h *UserHandler) Index(c *gin.Context) {
	var users []models.User

	if err := h.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	results := make([]UserResponse, 0, len(users))

	for _, p := range users {
		results = append(results, UserResponse{
			User: p,
			Link: config.App.Url + "/api/user/" + strconv.FormatUint(uint64(p.ID), 10),
		})
	}

	c.JSON(http.StatusOK, results)
}

func (h *UserHandler) Show(c *gin.Context) {
	id := c.Param("id")
	var p models.User

	if err := h.DB.First(&p, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, p)
}

func (h *UserHandler) Create(c *gin.Context) {
	var payload models.User

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	payload.Password = string(hash)

	if err := h.DB.Create(&payload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, payload)
}

func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	var payload models.User

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	payload.Password = string(hash)

	if err := h.DB.Model(&user).Updates(payload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Destroy(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	if err := h.DB.Delete(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	c.Status(http.StatusNoContent)
}
