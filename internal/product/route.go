package product

import "github.com/gin-gonic/gin"


func RegisterRoutes(
	router *gin.RouterGroup,
	handler *Handler,
	requireAuth gin.HandlerFunc,
	requireRole func(...string) gin.HandlerFunc,
) {
	productRoutes := router.Group("/products")

	productRoutes.GET(
		"",
		handler.GetAll,
	)

	productRoutes.GET(
		"/:id",
		handler.GetByID,
	)

	productRoutes.POST(
		"",
		requireAuth,
		requireRole("admin", "super_admin"),
		handler.Create,
	)
	
	productRoutes.PUT(
		"/:id",
		requireAuth,
		requireRole("admin", "super_admin"),
		handler.Update,
	)

	productRoutes.DELETE(
		"/:id",
		requireAuth,
		requireRole("super_admin"),
		handler.Delete,
	)

}