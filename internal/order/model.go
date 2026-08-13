package order

import "time"

const (
	StatusPending   = "pending"
	StatusConfirmed = "confirmed"
	StatusShipped   = "shipped"
	StatusDelivered = "delivered"
	StatusCancelled = "cancelled"
)

type Order struct {
	ID              uint        `gorm:"primaryKey" json:"id"`
	UserID          uint        `gorm:"not null;index" json:"userId"`
	Status          string      `gorm:"not null;default:pending" json:"status"`
	TotalAmount     float64     `gorm:"type:numeric(12,2);not null" json:"totalAmount"`
	ShippingAddress string      `gorm:"not null" json:"shippingAddress"`
	Items           []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;" json:"items"`
	CreatedAt       time.Time   `json:"createdAt"`
	UpdatedAt       time.Time   `json:"updatedAt"`
}

type OrderItem struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	OrderID     uint      `gorm:"not null;index" json:"orderId"`
	ProductID   uint      `gorm:"not null" json:"productId"`
	ProductName string    `gorm:"not null" json:"productName"`
	UnitPrice   float64   `gorm:"type:numeric(12,2);not null" json:"unitPrice"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	Subtotal    float64   `gorm:"type:numeric(12,2);not null" json:"subtotal"`
	CreatedAt   time.Time `json:"createdAt"`
}