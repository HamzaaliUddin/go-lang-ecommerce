package inventory

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

func (r *Repository) FindAll() ([]Inventory, error) {
	var inventories []Inventory

	err := r.db.
		Order("id DESC").
		Find(&inventories).
		Error

	if err != nil {
		return nil, err
	}

	return inventories, nil
}

func (r *Repository) FindByID(id uint) (*Inventory, error) {
	var inventory Inventory

	err := r.db.First(&inventory, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &inventory, nil
}

func (r *Repository) FindByProductID(
	productID uint,
) (*Inventory, error) {
	var inventory Inventory

	err := r.db.
		Where("product_id = ?", productID).
		First(&inventory).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &inventory, nil
}

func (r *Repository) Create(
	inventory *Inventory,
) error {
	return r.db.Create(inventory).Error
}

func (r *Repository) Update(
	inventory *Inventory,
) error {
	return r.db.
		Model(inventory).
		Updates(map[string]any{
			"stock":               inventory.Stock,
			"low_stock_threshold": inventory.LowStockThreshold,
		}).
		Error
}

func (r *Repository) DeleteByID(id uint) error {
	return r.db.
		Delete(&Inventory{}, id).
		Error
}