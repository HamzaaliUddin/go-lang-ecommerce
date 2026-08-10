package main

import (
	"ecommerce-api/internal/auth"
	"ecommerce-api/internal/config"
	appMiddleware "ecommerce-api/internal/http/middleware"
	"ecommerce-api/internal/product"
	"ecommerce-api/internal/role"
	"ecommerce-api/internal/user"

	"gorm.io/gorm"
)

type AppDependencies struct {
	AuthHandler    *auth.Handler
	UserHandler    *user.Handler
	ProductHandler *product.Handler
	AuthMiddleware *appMiddleware.AuthMiddleware
	RoleMiddleware *appMiddleware.RoleMiddleware
}

func buildDependencies(
	db *gorm.DB,
	cfg config.Config,
) AppDependencies {

	// Repositories
	userRepository := user.NewRepository(db)
	roleRepository := role.NewRepository(db)
	productRepository := product.NewRepository(db)

	// Auth
	authService := auth.NewService(
		userRepository,
		roleRepository,
		cfg.JWTSecret,
	)

	authHandler := auth.NewHandler(authService)

	// User
	userService := user.NewService(
		userRepository,
	)
		authMiddleware := appMiddleware.NewAuth(
		cfg.JWTSecret,
	)

	roleMiddleware := appMiddleware.NewRole(
		userService,
	)

	userHandler := user.NewHandler(userService)

	// Product
	productService := product.NewService(
		productRepository,
	)

	productHandler := product.NewHandler(productService)

	return AppDependencies{
		AuthHandler: authHandler,
		UserHandler: userHandler,
		ProductHandler: productHandler,
		AuthMiddleware: authMiddleware,
		RoleMiddleware: roleMiddleware,
	}
}