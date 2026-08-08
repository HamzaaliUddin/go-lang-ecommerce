package seed

import (
	"fmt"

	"ecommerce-api/internal/permission"

	"gorm.io/gorm"
)

func seedPermissions(db *gorm.DB) error {
	permissions := []permission.Permission{
		{Name: "Create User", Slug: "user.create"},
		{Name: "Read User", Slug: "user.read"},
		{Name: "Update User", Slug: "user.update"},
		{Name: "Delete User", Slug: "user.delete"},

		{Name: "Create Role", Slug: "role.create"},
		{Name: "Read Role", Slug: "role.read"},
		{Name: "Update Role", Slug: "role.update"},
		{Name: "Delete Role", Slug: "role.delete"},

		{Name: "Create Permission", Slug: "permission.create"},
		{Name: "Read Permission", Slug: "permission.read"},
		{Name: "Update Permission", Slug: "permission.update"},
		{Name: "Delete Permission", Slug: "permission.delete"},
	}

	for _, item := range permissions {
		var existing permission.Permission

		result := db.Where(
			"slug = ?",
			item.Slug,
		).FirstOrCreate(&existing, item)

		if result.Error != nil {
			return fmt.Errorf(
				"failed to seed permission %s: %w",
				item.Slug,
				result.Error,
			)
		}
	}

	return nil
}