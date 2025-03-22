package ports

import "github.com/gaelzamora/spent-one/internal/domain"

type SpentRepository interface {
	Create(spent *domain.Spent) error
}