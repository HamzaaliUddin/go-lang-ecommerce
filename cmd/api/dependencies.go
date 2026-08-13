package main

import (
	"ecommerce-api/internal/auth"
	"ecommerce-api/internal/banner"
	"ecommerce-api/internal/cart"
	"ecommerce-api/internal/category"
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
	BannerHandler  *banner.Handler
	ProductHandler *product.Handler
	categoryHanler *category.Handler
	cartHandler    *cart.Handler
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
	categoryRepository := category.NewRepository(db)
	cartRepository := cart.NewRepository(db)

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


	// Category
	categoryService := category.NewService(categoryRepository)

	categoryHandler := category.NewHandler(categoryService)

	// Cart
	cartService := cart.NewService(cartRepository,productRepository)

	cartHandler := cart.NewHandler(cartService)


	return AppDependencies{
		AuthHandler: authHandler,
		UserHandler: userHandler,
		BannerHandler: bannerHandler,
		ProductHandler: productHandler,
		categoryHanler: categoryHandler,
		cartHandler: cartHandler,
		AuthMiddleware: authMiddleware,
		RoleMiddleware: roleMiddleware,
	}
}