package banner

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
	banners, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve banners"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"banners": banners})
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := parseBannerID(c.Param("id"))

	banner, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve banner"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"banner": banner})
}

func (h *Handler) Create(c *gin.Context) {
	var banner Banner

	if err := c.ShouldBindJSON(&banner); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "invalid request body",
		})
		return 
	}
	createdBanner, err := h.service.Create(&banner)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"banner": createdBanner,
	})

}

func (h *Handler) Update(c *gin.Context) {
	id, err := parseBannerID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid product id",
		})
		return
	}
	var banner Banner
	if err := c.ShouldBindJSON(&banner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	updatedBannner, err := h.service.Update(id, &banner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": updatedBannner,
	})

}

func (h *Handler) Delete(c *gin.Context) {
		id, err := parseBannerID(c.Param("id"))
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




func parseBannerID(value string) (uint, error) {
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