package order

import (
	"errors"

	"ecommerce-api/internal/cart"

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

func (r *Repository) FindAllByUser(userID uint) ([]Order, error) {
	var orders []Order

	err := r.db.
		Preload("Items").
		Where("user_id = ?", userID).
		Order("id DESC").
		Find(&orders).
		Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *Repository) FindByIDAndUser(
	id uint,
	userID uint,
) (*Order, error) {
	var order Order

	err := r.db.
		Preload("Items").
		Where("id = ? AND user_id = ?", id, userID).
		First(&order).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &order, nil
}

func (r *Repository) FindAll() ([]Order, error) {
	var orders []Order

	err := r.db.
		Preload("Items").
		Order("id DESC").
		Find(&orders).
		Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *Repository) FindByID(id uint) (*Order, error) {
	var order Order

	err := r.db.
		Preload("Items").
		First(&order, id).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &order, nil
}

func (r *Repository) CreateAndClearCart(
	order *Order,
	userID uint,
) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Items").Create(order).Error; err != nil {
			return err
		}

		for i := range order.Items {
			order.Items[i].OrderID = order.ID
		}

		if len(order.Items) > 0 {
			if err := tx.Create(&order.Items).Error; err != nil {
				return err
			}
		}

		if err := tx.
			Where("user_id = ?", userID).
			Delete(&cart.CartItem{}).
			Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *Repository) UpdateStatus(
	order *Order,
	status string,
) error {
	return r.db.
		Model(order).
		Update("status", status).
		Error
}