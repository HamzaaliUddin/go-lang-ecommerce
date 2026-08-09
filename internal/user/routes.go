package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.RouterGroup,
	handler *Handler,
	requireAuth gin.HandlerFunc,
	requireRole func(...string) gin.HandlerFunc,
) {
	userRoutes := router.Group("/users")

	userRoutes.GET(
		"/me",
		requireAuth,
		handler.GetProfile,
	)

	userRoutes.GET(
		"",
		requireAuth,
		requireRole(
			"admin",
			"super_admin",
		),
		handler.GetAll,
	)

	userRoutes.GET(
		"/:id",
		requireAuth,
		requireRole(
			"admin",
			"super_admin",
		),
		handler.GetByID,
	)

	userRoutes.DELETE(
		"/:id",
		requireAuth,
		requireRole("super_admin"),
		handler.Delete,
	)
}