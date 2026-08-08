package router

import (
	"net/http"

	appMiddleware "ecommerce-api/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
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

	return router
}