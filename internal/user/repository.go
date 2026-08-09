package user

import (
	"errors"
	"ecommerce-api/internal/role"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var user User

	err := r.db.
		Preload("Roles.Permissions").
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateWithRole(
	account *User,
	assignedRole role.Role,
) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(account).Error; err != nil {
			return err
		}

		if err := tx.Model(account).
			Association("Roles").
			Append(&assignedRole); err != nil {
			return err
		}

		return nil
	})
}