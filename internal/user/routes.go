package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.RouterGroup,
	handler *Handler,
) {
	userRoutes := router.Group("/users")

	userRoutes.GET("", handler.GetAll)

	userRoutes.GET("/me", handler.GetProfile)

	userRoutes.GET("/:id", handler.GetByID)

	userRoutes.DELETE("/:id", handler.Delete)
}