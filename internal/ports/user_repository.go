package ports

import "github.com/gaelzamora/spent-one/internal/domain"

type UserRepository interface {
	FindByUsername(username string) (*domain.User, error)
	FindByID(id uint) (*domain.User, error)
	Create(user *domain.User) error
}

type SpentRepository interface {
	Create(spent *domain.Spent) error
}
