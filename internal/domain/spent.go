package domain

import "time"

type Spent struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Amount    float64   `json:"amount"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UserID    uint      `json:"user_id"`
}
