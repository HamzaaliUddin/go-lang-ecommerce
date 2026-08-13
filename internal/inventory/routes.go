package inventory

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.RouterGroup,
	handler *Handler,
	requireAuth gin.HandlerFunc,
	requireRole func(...string) gin.HandlerFunc,
) {
	inventoryRoutes :=
		router.Group("/admin/inventory")

	inventoryRoutes.Use(requireAuth)

	inventoryRoutes.Use(
		requireRole(
			"admin",
			"super_admin",
		),
	)

	inventoryRoutes.GET(
		"",
		handler.GetAll,
	)

	inventoryRoutes.GET(
		"/:id",
		handler.GetByID,
	)

	inventoryRoutes.POST(
		"",
		handler.Create,
	)

	inventoryRoutes.PATCH(
		"/:id",
		handler.Update,
	)

	inventoryRoutes.DELETE(
		"/:id",
		handler.Delete,
	)
}