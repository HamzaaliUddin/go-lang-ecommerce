package payment

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

type CreatePaymentRequest struct {
	OrderID uint   `json:"orderId" binding:"required"`
	Method  string `json:"method" binding:"required"`
}

type UpdatePaymentStatusRequest struct {
	Status        string `json:"status" binding:"required"`
	TransactionID string `json:"transactionId"`
}

func (h *Handler) GetMyPayments(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	payments, err := h.service.GetMyPayments(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payments": payments,
	})
}

func (h *Handler) Create(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	var request CreatePaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	createdPayment, err := h.service.Create(
		userID,
		request.OrderID,
		request.Method,
	)

	if err != nil {
		switch {
		case errors.Is(err, ErrOrderNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"message": "order not found",
			})

		case errors.Is(err, ErrPaymentAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{
				"message": "payment already exists",
			})

		case errors.Is(err, ErrInvalidPaymentMethod):
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid payment method",
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
		}

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"payment": createdPayment,
	})
}

func (h *Handler) GetAll(c *gin.Context) {
	payments, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payments": payments,
	})
}

func (h *Handler) UpdateStatus(c *gin.Context) {
	id, err := parsePaymentID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid payment id",
		})
		return
	}

	var request UpdatePaymentStatusRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	updatedPayment, err := h.service.UpdateStatus(
		id,
		request.Status,
		request.TransactionID,
	)

	if err != nil {
		switch {
		case errors.Is(err, ErrPaymentNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"message": "payment not found",
			})

		case errors.Is(err, ErrInvalidPaymentStatus):
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid payment status",
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payment": updatedPayment,
	})
}

func getUserID(c *gin.Context) (uint, bool) {
	value, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return 0, false
	}

	userID, ok := value.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return 0, false
	}

	return userID, true
}

func parsePaymentID(value string) (uint, error) {
	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil || id == 0 {
		return 0, errors.New("invalid payment id")
	}

	return uint(id), nil
}