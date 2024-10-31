package models

import "time"

type Bid struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ItemID    uint      `json:"item_id" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Price     int       `json:"price" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	Item      Item      `json:"item,omitempty" gorm:"foreignKey:ItemID"`
	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
