package models

import (
	"time"
)

type Swipe struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	ProfileID uint   `gorm:"not null" json:"profile_id"`
	SwipeType string `gorm:"not null" json:"swipe_type"` // "like" or "pass"
	CreatedAt time.Time
}
