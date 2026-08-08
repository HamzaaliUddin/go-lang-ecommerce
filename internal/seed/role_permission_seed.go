package seed

import (
	"fmt"

	"ecommerce-api/internal/permission"
	"ecommerce-api/internal/role"
	
	"gorm.io/gorm"
)

func assignRolePermissions(db *gorm.DB) error {
	if err := assignPermissionsToRole(db, "super_admin", []string{
		"user.create",
		"user.read",
		"user.update",
		"user.delete",
		"role.create",
		"role.read",
		"role.update",
		"role.delete",
		"permission.create",
		"permission.read",
		"permission.update",
		"permission.delete",
	}); err != nil {
		return err
	}

	if err := assignPermissionsToRole(db, "admin", []string{
		"user.create",
		"user.read",
		"user.update",
		"role.read",
		"permission.read",
	}); err != nil {
		return err
	}

	if err := assignPermissionsToRole(db, "customer", []string{
		"user.read",
	}); err != nil {
		return err
	}

	return nil
}

func assignPermissionsToRole(
	db *gorm.DB,
	roleSlug string,
	permissionSlugs []string,
) error {
	var selectedRole role.Role

	if err := db.Where("slug = ?", roleSlug).
		First(&selectedRole).Error; err != nil {
		return fmt.Errorf(
			"failed to find role %s: %w",
			roleSlug,
			err,
		)
	}

	var permissions []permission.Permission

	if err := db.Where("slug IN ?", permissionSlugs).
		Find(&permissions).Error; err != nil {
		return fmt.Errorf(
			"failed to find permissions for role %s: %w",
			roleSlug,
			err,
		)
	}

	if err := db.Model(&selectedRole).
		Association("Permissions").
		Replace(permissions); err != nil {
		return fmt.Errorf(
			"failed to assign permissions to role %s: %w",
			roleSlug,
			err,
		)
	}

	return nil
}