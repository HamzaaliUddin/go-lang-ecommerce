package cart

import "errors"

var (
	ErrCartItemNotFound = errors.New("cart item not found")
	ErrProductNotFound  = errors.New("product not found")
	ErrInvalidQuantity  = errors.New("quantity must be greater than zero")
)