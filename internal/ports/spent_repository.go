package ports

import "github.com/gaelzamora/spent-one/internal/domain"

type SpentRepository interface {
	Create(spent *domain.Spent) error
	GetSpents(userID uint) ([]domain.Spent, error)
	GetSpent(userID uint, spentID uint) (domain.Spent, error)
	DeleteSpent(userID uint, spentID uint) error
	UpdateSpent(spent domain.Spent) error
}
