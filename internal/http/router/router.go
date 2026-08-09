package router

import (
	"ecommerce-api/internal/auth"
	appMiddleware "ecommerce-api/internal/http/middleware"
	"ecommerce-api/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(authHandler *auth.Handler,userHandler *user.Handler,) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(appMiddleware.Security())
	router.Use(appMiddleware.CORS())

	api := router.Group("/api/v1")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "E-commerce API is running",
		})
	})

	auth.RegisterRoutes(api, authHandler)
	user.RegisterRoutes(api,userHandler)
	return router
}