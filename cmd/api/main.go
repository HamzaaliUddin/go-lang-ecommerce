package main

import (
	"ecommerce-api/internal/config"
	"ecommerce-api/internal/http/router"
	"ecommerce-api/internal/permission"
	"ecommerce-api/internal/platform/database"
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
		&user.User{}); err != nil {
		log.Fatal("database migration failed: ", err)
	}

	log.Println("database models synchronized successfully")
	if err := seed.Run(db, cfg); err != nil {
		log.Fatal("database seed failed: ", err)
	}
	log.Println("database seeded successfully")
	
	app := router.New()

	if err := app.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}