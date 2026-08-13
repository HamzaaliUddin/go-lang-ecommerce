package order

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.RouterGroup,
	handler *Handler,
	requireAuth gin.HandlerFunc,
	requireRole func(...string) gin.HandlerFunc,
) {
	orderRoutes := router.Group("/orders")

	orderRoutes.Use(requireAuth)

	orderRoutes.GET(
		"",
		handler.GetMyOrders,
	)

	orderRoutes.GET(
		"/:id",
		handler.GetMyOrder,
	)

	orderRoutes.POST(
		"",
		handler.Create,
	)

	adminRoutes := router.Group("/admin/orders")

	adminRoutes.Use(requireAuth)
	adminRoutes.Use(
		requireRole(
			"admin",
			"super_admin",
		),
	)

	adminRoutes.GET(
		"",
		handler.GetAll,
	)

	adminRoutes.GET(
		"/:id",
		handler.GetByID,
	)

	adminRoutes.PATCH(
		"/:id/status",
		handler.UpdateStatus,
	)
}