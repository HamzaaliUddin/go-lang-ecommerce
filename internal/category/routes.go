package category

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.RouterGroup,
	handler *Handler,
	requireAuth gin.HandlerFunc,
	requireRole func(...string) gin.HandlerFunc,
) {
	categoryRoutes := router.Group("/categories")

	categoryRoutes.GET("", handler.GetAll)

	categoryRoutes.GET("/:id", handler.GetByID)

	categoryRoutes.POST(
		"",
		requireAuth,
		requireRole("admin", "super_admin"),
		handler.Create,
	)

	categoryRoutes.PATCH(
		"/:id",
		requireAuth,
		requireRole("admin", "super_admin"),
		handler.Update,
	)

	categoryRoutes.DELETE(
		"/:id",
		requireAuth,
		requireRole("super_admin"),
		handler.Delete,
	)
}