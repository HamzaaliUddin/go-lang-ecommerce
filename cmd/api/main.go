package main

import (
	"log"

	"ecommerce-api/internal/http/router"
)

func main() {
	app := router.New()

	err := app.Run(":8080")
	if err != nil {
		log.Fatal("failed to start server:", err)
	}
}