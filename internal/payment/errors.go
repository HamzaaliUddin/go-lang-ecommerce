package payment

import "errors"

var (
	ErrPaymentNotFound      = errors.New("payment not found")
	ErrOrderNotFound        = errors.New("order not found")
	ErrPaymentAlreadyExists = errors.New("payment already exists")
	ErrInvalidPaymentMethod = errors.New("invalid payment method")
	ErrInvalidPaymentStatus = errors.New("invalid payment status")
)