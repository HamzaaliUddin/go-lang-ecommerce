package role

import (
	"time"

	"ecommerce-api/internal/permission"
)

type Role struct {
	ID          uint                    `gorm:"primaryKey" json:"id"`
	Name        string                  `gorm:"uniqueIndex;not null" json:"name"`
	Slug        string                  `gorm:"uniqueIndex;not null" json:"slug"`
	Description string                  `json:"description"`
	Permissions []permission.Permission `gorm:"many2many:role_permissions;" json:"permissions"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}