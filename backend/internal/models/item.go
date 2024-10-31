package models

import "time"

type Item struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Title        string    `json:"title" gorm:"not null;size:100"`
	Description  string    `json:"description" gorm:"type:text"`
	StartPrice   int       `json:"start_price" gorm:"not null"`
	CurrentPrice int       `json:"current_price" gorm:"not null"`
	StartAt      time.Time `json:"start_at" gorm:"not null"`
	EndAt        time.Time `json:"end_at" gorm:"not null"`
	Status       string    `json:"status" gorm:"not null;size:20"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null"`
	Bids         []Bid     `json:"bids,omitempty" gorm:"foreignKey:ItemID"`
}
