package category

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
	categories, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := parseCategoryID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid category id",
		})
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "category not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

func (h *Handler) Create(c *gin.Context) {
	var category Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	createdCategory, err := h.service.Create(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"category": createdCategory,
	})
}

func (h *Handler) Update(c *gin.Context) {
	id, err := parseCategoryID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid category id",
		})
		return
	}

	var category Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	updatedCategory, err := h.service.Update(id, &category)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "category not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": updatedCategory,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := parseCategoryID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid category id",
		})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "category not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.Status(http.StatusNoContent)
}
func parseCategoryID(value string) (uint, error) {
	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil || id == 0 {
		return 0, errors.New("invalid category id")
	}

	return uint(id), nil
}