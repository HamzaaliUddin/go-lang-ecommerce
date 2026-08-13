package order

import (
	"ecommerce-api/internal/cart"
)

type Service struct {
	repository     *Repository
	cartRepository *cart.Repository
}

func NewService(
	repository *Repository,
	cartRepository *cart.Repository,
) *Service {
	return &Service{
		repository:     repository,
		cartRepository: cartRepository,
	}
}

func (s *Service) GetMyOrders(
	userID uint,
) ([]Order, error) {
	return s.repository.FindAllByUser(userID)
}

func (s *Service) GetMyOrder(
	id uint,
	userID uint,
) (*Order, error) {
	foundOrder, err := s.repository.FindByIDAndUser(
		id,
		userID,
	)
	if err != nil {
		return nil, err
	}

	if foundOrder == nil {
		return nil, ErrOrderNotFound
	}

	return foundOrder, nil
}

func (s *Service) GetAll() ([]Order, error) {
	return s.repository.FindAll()
}

func (s *Service) GetByID(id uint) (*Order, error) {
	foundOrder, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if foundOrder == nil {
		return nil, ErrOrderNotFound
	}

	return foundOrder, nil
}

func (s *Service) CreateFromCart(
	userID uint,
	shippingAddress string,
) (*Order, error) {
	cartItems, err := s.cartRepository.FindAllByUser(userID)
	if err != nil {
		return nil, err
	}

	if len(cartItems) == 0 {
		return nil, ErrCartEmpty
	}

	orderItems := make(
		[]OrderItem,
		0,
		len(cartItems),
	)

	var totalAmount float64

	for _, cartItem := range cartItems {
		subtotal :=
			cartItem.Product.Price *
				float64(cartItem.Quantity)

		totalAmount += subtotal

		orderItems = append(
			orderItems,
			OrderItem{
				ProductID:   cartItem.ProductID,
				ProductName: cartItem.Product.Name,
				UnitPrice:   cartItem.Product.Price,
				Quantity:    cartItem.Quantity,
				Subtotal:    subtotal,
			},
		)
	}

	newOrder := &Order{
		UserID:          userID,
		Status:          StatusPending,
		TotalAmount:     totalAmount,
		ShippingAddress: shippingAddress,
		Items:           orderItems,
	}

	if err := s.repository.CreateAndClearCart(
		newOrder,
		userID,
	); err != nil {
		return nil, err
	}

	return s.repository.FindByIDAndUser(
		newOrder.ID,
		userID,
	)
}

func (s *Service) UpdateStatus(
	id uint,
	status string,
) (*Order, error) {
	if !isValidStatus(status) {
		return nil, ErrInvalidOrderStatus
	}

	foundOrder, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if foundOrder == nil {
		return nil, ErrOrderNotFound
	}

	if err := s.repository.UpdateStatus(
		foundOrder,
		status,
	); err != nil {
		return nil, err
	}

	return s.repository.FindByID(id)
}

func isValidStatus(status string) bool {
	switch status {
	case StatusPending,
		StatusConfirmed,
		StatusShipped,
		StatusDelivered,
		StatusCancelled:
		return true
	default:
		return false
	}
}