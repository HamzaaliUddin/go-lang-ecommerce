package category

import (
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}


func NewRepository(db *gorm.DB) *Repository{
	return &Repository{
		db:db,
	}
}


func (r *Repository) FindAll() ([]CategoryResponse,error){
	var category []CategoryResponse

	err := r.db.Find(&category).Error

	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil, nil
		}
		return nil,err
	}

	return category,nil
}

func (r *Repository) FindByID(id uint)(*Category,error){
	var category Category

	err := r.db.First(&category,id).Error

	if err != nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil , nil
		}
		return nil,err
	}

	return  &category,nil
	
}


func (r *Repository) Create(category Category) (*Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *Repository) Update(category *Category) (*Category, error) {
	err := r.db.Save(category).Error
	if err != nil {
		return nil, err
	}

	return category, nil
}
func (r *Repository) DeleteByID(id uint) error {
	err := r.db.Delete(&Category{}, id).Error
	if err != nil {
		return err
	}

	return nil
}