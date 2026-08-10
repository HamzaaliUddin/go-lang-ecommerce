package product

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetAll(c *gin.Context) {
	products, err := h.service.GetAll()

	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "products not found",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := parseProductID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid product id",
		})
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "product not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func (h *Handler) Create(c *gin.Context) {
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	createdProduct, err := h.service.Create(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"product": createdProduct,
	})
}

func (h *Handler) Update(c *gin.Context) {
	id, err := parseProductID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid product id",
		})
		return
	}
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	updatedProduct, err := h.service.Update(id, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": updatedProduct,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := parseProductID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid product id",
		})
		return
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "product not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "product deleted successfully",
	})
}




func parseProductID(value string) (uint, error) {
	id, err := strconv.ParseUint(
		value,
		10,
		64,
	)

	if err != nil {
		return 0, err
	}

	return uint(id), nil
}