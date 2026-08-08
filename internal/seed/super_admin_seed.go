package seed

import (
	"errors"
	"fmt"

	"ecommerce-api/internal/config"
	"ecommerce-api/internal/role"
	"ecommerce-api/internal/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func seedSuperAdmin(db *gorm.DB, cfg config.Config) error {
	var existingUser user.User

	result := db.Where(
		"email = ?",
		cfg.SuperAdminEmail,
	).First(&existingUser)

	if result.Error == nil {
		return nil
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf(
			"failed to check super admin: %w",
			result.Error,
		)
	}

	var superAdminRole role.Role

	if err := db.Where(
		"slug = ?",
		"super_admin",
	).First(&superAdminRole).Error; err != nil {
		return fmt.Errorf(
			"failed to find super admin role: %w",
			err,
		)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(cfg.SuperAdminPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return fmt.Errorf(
			"failed to hash super admin password: %w",
			err,
		)
	}

	superAdmin := user.User{
		FirstName:    cfg.SuperAdminFirstName,
		LastName:     cfg.SuperAdminLastName,
		Email:        cfg.SuperAdminEmail,
		PasswordHash: string(hashedPassword),
		IsActive:     true,
		Roles: []role.Role{
			superAdminRole,
		},
	}

	if err := db.Create(&superAdmin).Error; err != nil {
		return fmt.Errorf(
			"failed to create super admin: %w",
			err,
		)
	}

	return nil
}