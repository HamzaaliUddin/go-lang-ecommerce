package user

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
	users, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := parseUserID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}

	account, err := h.service.GetByID(id)

	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *Handler) GetProfile(c *gin.Context) {
	value, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	userID, ok := value.(uint)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	account, err := h.service.GetProfile(userID)

	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := parseUserID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
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

func parseUserID(value string) (uint, error) {
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