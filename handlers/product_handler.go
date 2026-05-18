package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nazala-mafa/go-crud/config"
	"github.com/nazala-mafa/go-crud/models"
	"gorm.io/gorm"
)

type ProductHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewProductHandler(db *gorm.DB, cfg *config.Config) *ProductHandler {
	return &ProductHandler{
		db, cfg,
	}
}

type ProductResponse struct {
	models.Product
	Link string `json:"link"`
}

func (h *ProductHandler) Index(c *gin.Context) {
	var products []models.Product

	if err := h.db.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	results := make([]ProductResponse, 0, len(products))

	for _, p := range products {
		results = append(results, ProductResponse{
			Product: p,
			Link:    h.cfg.App.Url + "/api/product/" + strconv.FormatUint(uint64(p.ID), 10),
		})
	}

	c.JSON(http.StatusOK, results)
}

func (h *ProductHandler) Show(c *gin.Context) {
	id := c.Param("id")
	var p models.Product

	if err := h.db.First(&p, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, p)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var payload models.Product

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.db.Create(&payload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, payload)
}

func (h *ProductHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if err := h.db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	var payload models.Product

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.db.Model(&product).Updates(payload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Destroy(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if err := h.db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	if err := h.db.Delete(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	c.Status(http.StatusNoContent)
}
