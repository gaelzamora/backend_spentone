package domain

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"-"`
}
