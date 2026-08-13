package main

import (
	"ecommerce-api/internal/banner"
	"ecommerce-api/internal/cart"
	"ecommerce-api/internal/category"
	"ecommerce-api/internal/config"
	"ecommerce-api/internal/http/router"
	"ecommerce-api/internal/inventory"
	"ecommerce-api/internal/order"
	"ecommerce-api/internal/payment"
	"ecommerce-api/internal/permission"
	"ecommerce-api/internal/platform/database"
	"ecommerce-api/internal/product"
	"ecommerce-api/internal/role"
	"ecommerce-api/internal/seed"
	"ecommerce-api/internal/user"
	"log"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get database instance: ", err)
	}

	defer sqlDB.Close()

	log.Println("database connected successfully")

	if err := db.AutoMigrate(
		&permission.Permission{},
		&role.Role{},
		&user.User{},
		&category.Category{},
		&product.Product{},
		&cart.CartItem{},
		&banner.Banner{},
		&order.Order{},
	  &order.OrderItem{},
		&payment.Payment{},
		&inventory.Inventory{},
		); err != nil {
		log.Fatal("database migration failed: ", err)
	}

	log.Println("database models synchronized successfully")
	if err := seed.Run(db, cfg); err != nil {
		log.Fatal("database seed failed: ", err)
	}
	log.Println("database seeded successfully")
	
	dependencies := buildDependencies(db, cfg)

	app := router.New(
		dependencies.AuthHandler,
		dependencies.UserHandler,
		dependencies.ProductHandler,
		dependencies.BannerHandler,
		dependencies.categoryHanler,
		dependencies.cartHandler,
		dependencies.OrderHandler,
		dependencies.PaymentHandler,
		dependencies.InventoryHandler,
		dependencies.AuthMiddleware,
		dependencies.RoleMiddleware,
	)

	if err := app.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}