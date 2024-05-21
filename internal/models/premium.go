package models

import (
	"time"
)

type PremiumPackage struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	PackageType string `gorm:"not null"` // e.g., "no_swipe_quota", "verified"
	ExpiryDate  time.Time
}
