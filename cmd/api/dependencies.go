package main

import (
	"ecommerce-api/internal/auth"
	"ecommerce-api/internal/config"
	appMiddleware "ecommerce-api/internal/http/middleware"
	"ecommerce-api/internal/product"
	"ecommerce-api/internal/role"
	"ecommerce-api/internal/user"
	"ecommerce-api/internal/banner"
	"gorm.io/gorm"
)

type AppDependencies struct {
	AuthHandler    *auth.Handler
	UserHandler    *user.Handler
	BannerHandler  *banner.Handler
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
	bannerRepository := banner.NewRepository(db)

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

	// Banner
	bannerService := banner.NewService(
		bannerRepository,
	)

	bannerHandler := banner.NewHandler(bannerService)

	return AppDependencies{
		AuthHandler: authHandler,
		UserHandler: userHandler,
		BannerHandler: bannerHandler,
		ProductHandler: productHandler,
		AuthMiddleware: authMiddleware,
		RoleMiddleware: roleMiddleware,
	}
}