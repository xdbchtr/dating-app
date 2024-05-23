package models

import (
	"time"
)

type PremiumPackage struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	PackageType string `gorm:"not null" json:"package_type"` // e.g., "no_swipe_quota", "verified"
	ExpiryDate  time.Time
}
