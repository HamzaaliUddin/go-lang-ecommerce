package cart

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.RouterGroup,
	handler *Handler,
	requireAuth gin.HandlerFunc,
) {
	cartRoutes := router.Group("/cart")

	cartRoutes.Use(requireAuth)

	cartRoutes.GET("", handler.GetAll)

	cartRoutes.POST(
		"/items",
		handler.AddItem,
	)

	cartRoutes.PATCH(
		"/items/:id",
		handler.UpdateQuantity,
	)

	cartRoutes.DELETE(
		"/items/:id",
		handler.Delete,
	)

	cartRoutes.DELETE(
		"",
		handler.Clear,
	)
}