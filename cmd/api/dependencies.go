package main

import (
	"ecommerce-api/internal/auth"
	"ecommerce-api/internal/config"
	"ecommerce-api/internal/role"
	"ecommerce-api/internal/user"

	"gorm.io/gorm"
)

type AppDependencies struct {
	AuthHandler *auth.Handler
	UserHandler *user.Handler
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

	userHandler := user.NewHandler(userService)

	return AppDependencies{
		AuthHandler: authHandler,
		UserHandler: userHandler,
	}
}