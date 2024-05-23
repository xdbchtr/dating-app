package models

import (
	"time"
)

type Swipe struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	ProfileID uint      `gorm:"not null" json:"profile_id"`
	SwipeType string    `gorm:"not null" json:"swipe_type"`          // "like" or "pass"
	SwipedAt  time.Time `gorm:"not null;type:date" json:"swiped_at"` // Timestamp when the swipe occurred
}
