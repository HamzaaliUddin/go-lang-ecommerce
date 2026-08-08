package user

import (
	"time"

	"ecommerce-api/internal/role"
)

type User struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	FirstName    string      `gorm:"not null" json:"firstName"`
	LastName     string      `gorm:"not null" json:"lastName"`
	Email        string      `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string      `gorm:"not null" json:"-"`
	IsActive     bool        `gorm:"default:true" json:"isActive"`
	Roles        []role.Role `gorm:"many2many:user_roles;" json:"roles"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}