package models

type Profile struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   uint   `gorm:"not null"`
	Name     string `gorm:"not null"`
	Age      int    `gorm:"not null"`
	Bio      string
	PhotoURL string `json:"photo_url"`
}
