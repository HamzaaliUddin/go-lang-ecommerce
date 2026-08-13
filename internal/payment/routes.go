package payment

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.RouterGroup,
	handler *Handler,
	requireAuth gin.HandlerFunc,
	requireRole func(...string) gin.HandlerFunc,
) {
	paymentRoutes := router.Group("/payments")
	paymentRoutes.Use(requireAuth)

	paymentRoutes.GET(
		"",
		handler.GetMyPayments,
	)

	paymentRoutes.POST(
		"",
		handler.Create,
	)

	adminRoutes := router.Group("/admin/payments")

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

	adminRoutes.PATCH(
		"/:id/status",
		handler.UpdateStatus,
	)
}