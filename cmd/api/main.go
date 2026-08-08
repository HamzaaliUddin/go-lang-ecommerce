package main

import (
	"log"

	"ecommerce-api/internal/http/router"
	"ecommerce-api/internal/config"
)

func main() {
	cfg := config.Load()
	app := router.New()


	err := app.Run(":" + cfg.AppPort)
	if err != nil {
		log.Fatal("failed to start server:", err)
	}
}