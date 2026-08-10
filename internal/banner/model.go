package banner

import (
	"time"
)

type Banner struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ImageURL  string `gorm:"not null" json:"imageUrl"`
	LinkURL   string `json:"linkUrl"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}