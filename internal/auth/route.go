package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.RouterGroup,
	handler *Handler,
) {
	authRoutes := router.Group("/auth")

	authRoutes.POST("/login",handler.Login)
	authRoutes.POST("/sign-up", handler.SignUp)
}