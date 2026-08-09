package seed

import (
	"errors"
	"fmt"

	"ecommerce-api/internal/role"
	"ecommerce-api/internal/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type customerSeedData struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func seedCustomers(db *gorm.DB) error {
	customers := []customerSeedData{
		{
			FirstName: "Ali",
			LastName:  "Khan",
			Email:     "ali@example.com",
			Password:  "Customer@123",
		},
		{
			FirstName: "Sara",
			LastName:  "Ahmed",
			Email:     "sara@example.com",
			Password:  "Customer@123",
		},
		{
			FirstName: "Usman",
			LastName:  "Raza",
			Email:     "usman@example.com",
			Password:  "Customer@123",
		},
	}

	var customerRole role.Role

	if err := db.Where(
		"slug = ?",
		"customer",
	).First(&customerRole).Error; err != nil {
		return fmt.Errorf(
			"failed to find customer role: %w",
			err,
		)
	}

	for _, item := range customers {
		if err := seedCustomer(
			db,
			item,
			customerRole,
		); err != nil {
			return err
		}
	}

	return nil
}

func seedCustomer(
	db *gorm.DB,
	data customerSeedData,
	customerRole role.Role,
) error {
	var existingUser user.User

	result := db.Where(
		"email = ?",
		data.Email,
	).First(&existingUser)

	if result.Error == nil {
		return nil
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf(
			"failed to check customer %s: %w",
			data.Email,
			result.Error,
		)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(data.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return fmt.Errorf(
			"failed to hash password for %s: %w",
			data.Email,
			err,
		)
	}

	customer := user.User{
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Email:        data.Email,
		PasswordHash: string(hashedPassword),
		IsActive:     true,
		Roles: []role.Role{
			customerRole,
		},
	}

	if err := db.Create(&customer).Error; err != nil {
		return fmt.Errorf(
			"failed to create customer %s: %w",
			data.Email,
			err,
		)
	}

	return nil
}