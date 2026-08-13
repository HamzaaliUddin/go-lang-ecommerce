package cart

import (
	"ecommerce-api/internal/product"
)

type Service struct {
	repository        *Repository
	productRepository *product.Repository
}

func NewService(
	repository *Repository,
	productRepository *product.Repository,
) *Service {
	return &Service{
		repository:        repository,
		productRepository: productRepository,
	}
}

func (s *Service) GetAll(userID uint) ([]CartItem, error) {
	return s.repository.FindAllByUser(userID)
}

func (s *Service) AddItem(
	userID uint,
	productID uint,
	quantity int,
) (*CartItem, error) {
	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	foundProduct, err := s.productRepository.FindByID(productID)
	if err != nil {
		return nil, err
	}

	if foundProduct == nil {
		return nil, ErrProductNotFound
	}

	existingItem, err := s.repository.FindByUserAndProduct(
		userID,
		productID,
	)
	if err != nil {
		return nil, err
	}

	if existingItem != nil {
		existingItem.Quantity += quantity

		if err := s.repository.Update(existingItem); err != nil {
			return nil, err
		}

		return s.repository.FindByIDAndUser(
			existingItem.ID,
			userID,
		)
	}

	item := &CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}

	if err := s.repository.Create(item); err != nil {
		return nil, err
	}

	return s.repository.FindByIDAndUser(
		item.ID,
		userID,
	)
}

func (s *Service) UpdateQuantity(
	userID uint,
	itemID uint,
	quantity int,
) (*CartItem, error) {
	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	item, err := s.repository.FindByIDAndUser(
		itemID,
		userID,
	)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, ErrCartItemNotFound
	}

	item.Quantity = quantity

	if err := s.repository.Update(item); err != nil {
		return nil, err
	}

	return s.repository.FindByIDAndUser(
		item.ID,
		userID,
	)
}

func (s *Service) Delete(
	userID uint,
	itemID uint,
) error {
	item, err := s.repository.FindByIDAndUser(
		itemID,
		userID,
	)
	if err != nil {
		return err
	}

	if item == nil {
		return ErrCartItemNotFound
	}

	return s.repository.DeleteByIDAndUser(
		itemID,
		userID,
	)
}

func (s *Service) Clear(userID uint) error {
	return s.repository.ClearByUser(userID)
}