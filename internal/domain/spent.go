package domain

type Spent struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Amount float64 `json:"amount"`
	Reason string `json:"reason"`
	SpentAt string `json:"spent_at"`
	UserID uint `json:"user_id"`
}