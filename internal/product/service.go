package product

import (
	"errors"
	"fmt"
)


var ErrUserNotFound = errors.New("user not found")

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetByID(id uint) (*ProductResponse, error) {
	product, err := s.repository.FindByID(id)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to find user: %w",
			err,
		)
	}

	if product == nil {
		return nil, ErrUserNotFound
	}

	response := toProductResponse(*product)
	return &response, nil
}

func (s *Service) GetAll() ([]ProductResponse, error) {
	products, err := s.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to fetch products: %w",
			err,
		)
	}
	if len(products) == 0 {
		return nil, ErrUserNotFound
	}

	response := make([]ProductResponse, 0, len(products))
	for _, product := range products {
		response = append(
			response,
			toProductResponse(product),
		)
	}

	return response, nil
}

func (s *Service) Create(product *Product) (*ProductResponse, error) {
	createdProduct, err := s.repository.Create(product)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create product: %w",
			err,
		)
	}
	response := toProductResponse(*createdProduct)

	return &response, nil
}

func (s *Service) Update(id uint, product *Product) (*ProductResponse, error) {
	updatedProduct, err := s.repository.Update(product)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to update product: %w",
			err,
		)
	}

	response := toProductResponse(*updatedProduct)

	return &response, nil
}

func (s *Service) Delete(id uint) error {
	product, err := s.repository.FindByID(id)
	if err != nil {
		return fmt.Errorf(
			"failed to find product: %w",
			err,
		)
	}
	if product == nil {
		return ErrUserNotFound
	}
	return s.repository.Delete(product)

}