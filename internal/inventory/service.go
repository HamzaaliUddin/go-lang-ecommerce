package inventory

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

func (s *Service) GetAll() ([]InventoryResponse, error) {
	inventories, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	response := make(
		[]InventoryResponse,
		0,
		len(inventories),
	)

	for _, inventory := range inventories {
		response = append(
			response,
			toInventoryResponse(inventory),
		)
	}

	return response, nil
}

func (s *Service) GetByID(
	id uint,
) (*Inventory, error) {
	foundInventory, err := s.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	if foundInventory == nil {
		return nil, ErrInventoryNotFound
	}

	return foundInventory, nil
}

func (s *Service) Create(
	request CreateInventoryRequest,
) (*InventoryResponse, error) {
	if request.Stock < 0 {
		return nil, ErrInvalidStock
	}

	foundProduct, err :=
		s.productRepository.FindByID(
			request.ProductID,
		)

	if err != nil {
		return nil, err
	}

	if foundProduct == nil {
		return nil, ErrProductNotFound
	}

	existingInventory, err :=
		s.repository.FindByProductID(
			request.ProductID,
		)

	if err != nil {
		return nil, err
	}

	if existingInventory != nil {
		return nil, ErrInventoryAlreadyExists
	}

	inventory := &Inventory{
		ProductID:         request.ProductID,
		Stock:             request.Stock,
		LowStockThreshold: request.LowStockThreshold,
	}

	if err := s.repository.Create(inventory); err != nil {
		return nil, err
	}

	response := toInventoryResponse(*inventory)

	return &response, nil
}

func (s *Service) Update(
	id uint,
	stock int,
	lowStockThreshold int,
) (*Inventory, error) {
	if stock < 0 {
		return nil, ErrInvalidStock
	}

	foundInventory, err :=
		s.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	if foundInventory == nil {
		return nil, ErrInventoryNotFound
	}

	foundInventory.Stock = stock
	foundInventory.LowStockThreshold = lowStockThreshold

	if err := s.repository.Update(
		foundInventory,
	); err != nil {
		return nil, err
	}

	return foundInventory, nil
}

func (s *Service) Delete(id uint) error {
	foundInventory, err :=
		s.repository.FindByID(id)

	if err != nil {
		return err
	}

	if foundInventory == nil {
		return ErrInventoryNotFound
	}

	return s.repository.DeleteByID(id)
}