package main

import (
	"ecommerce-api/internal/auth"
	"ecommerce-api/internal/config"
	"ecommerce-api/internal/role"
	"ecommerce-api/internal/user"
	appMiddleware "ecommerce-api/internal/http/middleware"
	"gorm.io/gorm"
)

type AppDependencies struct {
	AuthHandler    *auth.Handler
	UserHandler    *user.Handler
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

	return AppDependencies{
		AuthHandler: authHandler,
		UserHandler: userHandler,
		AuthMiddleware: authMiddleware,
		RoleMiddleware: roleMiddleware,
	}
}