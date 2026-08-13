package payment

import (
	"ecommerce-api/internal/order"
)

type Service struct {
	repository      *Repository
	orderRepository *order.Repository
}

func NewService(
	repository *Repository,
	orderRepository *order.Repository,
) *Service {
	return &Service{
		repository:      repository,
		orderRepository: orderRepository,
	}
}

func (s *Service) GetMyPayments(
	userID uint,
) ([]Payment, error) {
	return s.repository.FindAllByUser(userID)
}

func (s *Service) GetAll() ([]Payment, error) {
	return s.repository.FindAll()
}

func (s *Service) Create(
	userID uint,
	orderID uint,
	method string,
) (*Payment, error) {
	if !isValidMethod(method) {
		return nil, ErrInvalidPaymentMethod
	}

	foundOrder, err := s.orderRepository.FindByIDAndUser(
		orderID,
		userID,
	)
	if err != nil {
		return nil, err
	}

	if foundOrder == nil {
		return nil, ErrOrderNotFound
	}

	existingPayment, err := s.repository.FindByOrderID(orderID)
	if err != nil {
		return nil, err
	}

	if existingPayment != nil {
		return nil, ErrPaymentAlreadyExists
	}

	newPayment := &Payment{
		OrderID: orderID,
		UserID:  userID,
		Amount:  foundOrder.TotalAmount,
		Method:  method,
		Status:  StatusPending,
	}

	if err := s.repository.Create(newPayment); err != nil {
		return nil, err
	}

	return newPayment, nil
}

func (s *Service) UpdateStatus(
	id uint,
	status string,
	transactionID string,
) (*Payment, error) {
	if !isValidStatus(status) {
		return nil, ErrInvalidPaymentStatus
	}

	foundPayment, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if foundPayment == nil {
		return nil, ErrPaymentNotFound
	}

	if err := s.repository.UpdateStatus(
		foundPayment,
		status,
		transactionID,
	); err != nil {
		return nil, err
	}

	return s.repository.FindByID(id)
}

func isValidMethod(method string) bool {
	switch method {
	case MethodCOD, MethodCard:
		return true
	default:
		return false
	}
}

func isValidStatus(status string) bool {
	switch status {
	case StatusPending,
		StatusPaid,
		StatusFailed:
		return true
	default:
		return false
	}
}