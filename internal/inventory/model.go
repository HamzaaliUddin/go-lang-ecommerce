package inventory

import "time"

type Inventory struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	ProductID         uint      `gorm:"uniqueIndex;not null" json:"productId"`
	Stock             int       `gorm:"not null;default:0" json:"stock"`
	LowStockThreshold int       `gorm:"not null;default:5" json:"lowStockThreshold"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}