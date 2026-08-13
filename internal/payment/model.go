package payment

import "time"

const (
	MethodCOD  = "cod"
	MethodCard = "card"

	StatusPending = "pending"
	StatusPaid    = "paid"
	StatusFailed  = "failed"
)

type Payment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	OrderID       uint      `gorm:"not null;uniqueIndex" json:"orderId"`
	UserID        uint      `gorm:"not null;index" json:"userId"`
	Amount        float64   `gorm:"type:numeric(12,2);not null" json:"amount"`
	Method        string    `gorm:"not null" json:"method"`
	Status        string    `gorm:"not null;default:pending" json:"status"`
	TransactionID string    `json:"transactionId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}