package cart

import (
	"errors"

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

func (r *Repository) FindAllByUser(userID uint) ([]CartItem, error) {
	var items []CartItem

	err := r.db.
		Preload("Product").
		Where("user_id = ?", userID).
		Order("id DESC").
		Find(&items).
		Error

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Repository) FindByIDAndUser(
	id uint,
	userID uint,
) (*CartItem, error) {
	var item CartItem

	err := r.db.
		Preload("Product").
		Where("id = ? AND user_id = ?", id, userID).
		First(&item).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &item, nil
}

func (r *Repository) FindByUserAndProduct(
	userID uint,
	productID uint,
) (*CartItem, error) {
	var item CartItem

	err := r.db.
		Where(
			"user_id = ? AND product_id = ?",
			userID,
			productID,
		).
		First(&item).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &item, nil
}

func (r *Repository) Create(item *CartItem) error {
	return r.db.Create(item).Error
}

func (r *Repository) Update(item *CartItem) error {
	return r.db.
		Model(item).
		Update("quantity", item.Quantity).
		Error
}

func (r *Repository) DeleteByIDAndUser(
	id uint,
	userID uint,
) error {
	return r.db.
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&CartItem{}).
		Error
}

func (r *Repository) ClearByUser(userID uint) error {
	return r.db.
		Where("user_id = ?", userID).
		Delete(&CartItem{}).
		Error
}