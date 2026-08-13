package order

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

type CreateOrderRequest struct {
	ShippingAddress string `json:"shippingAddress" binding:"required,min=5"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

func (h *Handler) GetMyOrders(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	orders, err := h.service.GetMyOrders(userID)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "internal server error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"orders": orders,
		},
	)
}

func (h *Handler) GetMyOrder(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	id, err := parseOrderID(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid order id",
			},
		)
		return
	}

	foundOrder, err := h.service.GetMyOrder(
		id,
		userID,
	)
	if err != nil {
		if errors.Is(err, ErrOrderNotFound) {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"message": "order not found",
				},
			)
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "internal server error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"order": foundOrder,
		},
	)
}

func (h *Handler) Create(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	var request CreateOrderRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid request body",
			},
		)
		return
	}

	createdOrder, err := h.service.CreateFromCart(
		userID,
		request.ShippingAddress,
	)
	if err != nil {
		if errors.Is(err, ErrCartEmpty) {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"message": "cart is empty",
				},
			)
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "internal server error",
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"order": createdOrder,
		},
	)
}

func (h *Handler) GetAll(c *gin.Context) {
	orders, err := h.service.GetAll()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "internal server error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"orders": orders,
		},
	)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := parseOrderID(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid order id",
			},
		)
		return
	}

	foundOrder, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, ErrOrderNotFound) {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"message": "order not found",
				},
			)
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "internal server error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"order": foundOrder,
		},
	)
}

func (h *Handler) UpdateStatus(c *gin.Context) {
	id, err := parseOrderID(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid order id",
			},
		)
		return
	}

	var request UpdateOrderStatusRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid request body",
			},
		)
		return
	}

	updatedOrder, err := h.service.UpdateStatus(
		id,
		request.Status,
	)
	if err != nil {
		if errors.Is(err, ErrOrderNotFound) {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"message": "order not found",
				},
			)
			return
		}

		if errors.Is(err, ErrInvalidOrderStatus) {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"message": "invalid order status",
				},
			)
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "internal server error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"order": updatedOrder,
		},
	)
}

func getUserID(c *gin.Context) (uint, bool) {
	value, exists := c.Get("userID")
	if !exists {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "unauthorized",
			},
		)
		return 0, false
	}

	userID, ok := value.(uint)
	if !ok {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "unauthorized",
			},
		)
		return 0, false
	}

	return userID, true
}

func parseOrderID(value string) (uint, error) {
	id, err := strconv.ParseUint(
		value,
		10,
		64,
	)

	if err != nil || id == 0 {
		return 0, errors.New("invalid order id")
	}

	return uint(id), nil
}