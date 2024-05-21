package models

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	Email      string `gorm:"unique;not null"`
	IsVerified bool   `gorm:"default:false"`
}
