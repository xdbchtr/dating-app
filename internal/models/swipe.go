package models

import (
	"time"
)

type Swipe struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	ProfileID uint   `gorm:"not null"`
	SwipeType string `gorm:"not null"` // "like" or "pass"
	CreatedAt time.Time
}
