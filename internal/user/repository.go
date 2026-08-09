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

func (r *Repository) FindByID(id uint) (*User, error) {
	var account User

	err := r.db.
		Preload("Roles.Permissions").
		First(&account, id).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &account, nil
}

func (r *Repository) FindAll() ([]User, error) {
	var users []User

	err := r.db.
		Preload("Roles").
		Order("id DESC").
		Find(&users).
		Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) Delete(account *User) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(account).
			Association("Roles").
			Clear(); err != nil {
			return err
		}

		if err := tx.Delete(account).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *Repository) HasAnyRole(
	userID uint,
	roleSlugs ...string,
) (bool, error) {
	var count int64

	err := r.db.
		Table("roles").
		Joins(
			"JOIN user_roles ON user_roles.role_id = roles.id",
		).
		Where(
			"user_roles.user_id = ? AND roles.slug IN ?",
			userID,
			roleSlugs,
		).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}