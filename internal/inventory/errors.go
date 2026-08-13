package inventory

import "errors"

var (
	ErrInventoryNotFound      = errors.New("inventory not found")
	ErrProductNotFound        = errors.New("product not found")
	ErrInventoryAlreadyExists = errors.New("inventory already exists")
	ErrInvalidStock           = errors.New("stock cannot be negative")
)