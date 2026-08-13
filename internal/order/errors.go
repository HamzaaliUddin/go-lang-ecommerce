package order

import "errors"

var (
	ErrOrderNotFound     = errors.New("order not found")
	ErrCartEmpty         = errors.New("cart is empty")
	ErrInvalidOrderStatus = errors.New("invalid order status")
)