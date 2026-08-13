package payment

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

func (r *Repository) FindByID(id uint) (*Payment, error) {
	var payment Payment

	err := r.db.First(&payment, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &payment, nil
}

func (r *Repository) FindByOrderID(orderID uint) (*Payment, error) {
	var payment Payment

	err := r.db.
		Where("order_id = ?", orderID).
		First(&payment).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &payment, nil
}

func (r *Repository) FindAllByUser(userID uint) ([]Payment, error) {
	var payments []Payment

	err := r.db.
		Where("user_id = ?", userID).
		Order("id DESC").
		Find(&payments).
		Error

	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (r *Repository) FindAll() ([]Payment, error) {
	var payments []Payment

	err := r.db.
		Order("id DESC").
		Find(&payments).
		Error

	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (r *Repository) Create(payment *Payment) error {
	return r.db.Create(payment).Error
}

func (r *Repository) UpdateStatus(
	payment *Payment,
	status string,
	transactionID string,
) error {
	return r.db.
		Model(payment).
		Updates(map[string]any{
			"status":         status,
			"transaction_id": transactionID,
		}).
		Error
}