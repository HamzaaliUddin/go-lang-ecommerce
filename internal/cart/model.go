package cart

import (
	"time"

	"ecommerce-api/internal/product"
)

type CartItem struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	UserID    uint            `gorm:"not null;uniqueIndex:idx_user_product" json:"userId"`
	ProductID uint            `gorm:"not null;uniqueIndex:idx_user_product" json:"productId"`
	Product   product.Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int             `gorm:"not null;default:1" json:"quantity"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}