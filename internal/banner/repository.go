package banner

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

func (r *Repository) FindAll() ([]Banner, error) {
	var banners []Banner

	err := r.db.Find(&banners).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return banners, nil
}

func (r *Repository) FindByID(id uint) (*Banner, error) {
	var banner Banner
	err := r.db.First(&banner, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &banner, nil
}

func (r *Repository) Create(banner *Banner) (*Banner, error) {
	err := r.db.Create(banner).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil,err
		}
	}
	return banner, nil
}

func (r *Repository) Update(banner *Banner) (*Banner, error) {
	err := r.db.Save(banner).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	return banner, nil
}


func (r *Repository) Delete(banner *Banner) error {
	return r.db.Delete(banner).Error
}