package seed

import (
	"fmt"

	"ecommerce-api/internal/role"

	"gorm.io/gorm"
)

func seedRoles(db *gorm.DB) error {
	roles := []role.Role{
		{
			Name:        "Super Admin",
			Slug:        "super_admin",
			Description: "Full platform access",
		},
		{
			Name:        "Admin",
			Slug:        "admin",
			Description: "Platform administration access",
		},
		{
			Name:        "Customer",
			Slug:        "customer",
			Description: "Standard customer access",
		},
	}

	for _, item := range roles {
		var existing role.Role

		result := db.Where(
			"slug = ?",
			item.Slug,
		).FirstOrCreate(&existing, item)

		if result.Error != nil {
			return fmt.Errorf(
				"failed to seed role %s: %w",
				item.Slug,
				result.Error,
			)
		}
	}

	return nil
}