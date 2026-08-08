package seed

import (
	"ecommerce-api/internal/config"

	"gorm.io/gorm"
)

func Run(db *gorm.DB, cfg config.Config) error {
	if err := seedPermissions(db); err != nil {
		return err
	}

	if err := seedRoles(db); err != nil {
		return err
	}

	if err := assignRolePermissions(db); err != nil {
		return err
	}

	if err := seedSuperAdmin(db, cfg); err != nil {
		return err
	}

	return nil
}