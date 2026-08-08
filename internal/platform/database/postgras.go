package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(databaseURL),
		&gorm.Config{},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}