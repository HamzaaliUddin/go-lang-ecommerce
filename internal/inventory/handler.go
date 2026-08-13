package inventory

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
	inventories, err := h.service.GetAll()

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
			"inventories": inventories,
		},
	)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := parseInventoryID(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid inventory id",
			},
		)
		return
	}

	foundInventory, err := h.service.GetByID(id)

	if err != nil {
		if errors.Is(err, ErrInventoryNotFound) {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"message": "inventory not found",
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
			"inventory": foundInventory,
		},
	)
}

func (h *Handler) Create(c *gin.Context) {
	var request CreateInventoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	inventory, err := h.service.Create(request)
	if err != nil {
		// error handling
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"inventory": inventory,
	})
}

func (h *Handler) Update(c *gin.Context) {
	id, err := parseInventoryID(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid inventory id",
			},
		)
		return
	}

	var request UpdateInventoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid request body",
			},
		)
		return
	}

	updatedInventory, err := h.service.Update(
		id,
		request.Stock,
		request.LowStockThreshold,
	)

	if err != nil {
		switch {
		case errors.Is(err, ErrInventoryNotFound):
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"message": "inventory not found",
				},
			)

		case errors.Is(err, ErrInvalidStock):
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"message": "invalid stock",
				},
			)

		default:
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"message": "internal server error",
				},
			)
		}

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"inventory": updatedInventory,
		},
	)
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := parseInventoryID(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid inventory id",
			},
		)
		return
	}

	err = h.service.Delete(id)

	if err != nil {
		if errors.Is(err, ErrInventoryNotFound) {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"message": "inventory not found",
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

	c.Status(http.StatusNoContent)
}

func parseInventoryID(value string) (uint, error) {
	id, err := strconv.ParseUint(
		value,
		10,
		64,
	)

	if err != nil || id == 0 {
		return 0, errors.New(
			"invalid inventory id",
		)
	}

	return uint(id), nil
}