package router

import (
	"net/http"

	"ecommerce-api/internal/auth"
	"ecommerce-api/internal/banner"
	"ecommerce-api/internal/cart"
	"ecommerce-api/internal/category"
	appMiddleware "ecommerce-api/internal/http/middleware"
	"ecommerce-api/internal/order"
	"ecommerce-api/internal/product"
	"ecommerce-api/internal/user"

	"github.com/gin-gonic/gin"
)

func New(
	authHandler *auth.Handler,
	userHandler *user.Handler,
	productHandler *product.Handler,
	bannerHandler *banner.Handler,
	categoryHandler *category.Handler,
	cartHandler *cart.Handler,
	orderHandler *order.Handler,
	authMiddleware *appMiddleware.AuthMiddleware,
	roleMiddleware *appMiddleware.RoleMiddleware,
) *gin.Engine {
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

	auth.RegisterRoutes(
		api,
		authHandler,
	)

	user.RegisterRoutes(
		api,
		userHandler,
		authMiddleware.RequireAuth(),
		roleMiddleware.Require,
	)

	banner.RegisterRoutes(
		api,
		bannerHandler,
		authMiddleware.RequireAuth(),
		roleMiddleware.Require,
	)

	product.RegisterRoutes(
		api,
		productHandler,
		authMiddleware.RequireAuth(),
		roleMiddleware.Require,
	)

	category.RegisterRoutes(
	api,
	categoryHandler,
	authMiddleware.RequireAuth(),
	roleMiddleware.Require,
  )	

	cart.RegisterRoutes(
		api,
		cartHandler,
		authMiddleware.RequireAuth(),
	)

	return router
}