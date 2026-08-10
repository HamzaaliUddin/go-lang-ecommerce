package product

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

func (r *Repository) FindAll() ([]Product, error) {
	var products []Product

	err := r.db.Find(&products).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return products, err
}

func (r *Repository) FindByID(id uint) (*Product, error) {
	var product Product

	err := r.db.First(&product, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *Repository) Create(product *Product) (*Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *Repository) Update(product *Product) (*Product, error) {
	if err := r.db.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *Repository) Delete(product *Product) error {
	return r.db.Delete(product).Error	
}