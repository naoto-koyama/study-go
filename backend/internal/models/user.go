package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"not null;unique;size:100"`
	Password  string    `json:"password" gorm:"not null;size:100"`
	Username  string    `json:"username" gorm:"not null;size:100"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
}
