package application

import (
	"github.com/gaelzamora/spent-one/internal/domain"
	"github.com/gaelzamora/spent-one/internal/ports"
)

type SpentService struct {
	repo ports.SpentRepository
}

func NewSpentService(repo ports.SpentRepository) *SpentService {
	return &SpentService{repo: repo}
}

func (s *SpentService) CreateSpent(spent *domain.Spent) (domain.Spent, error) {
	err := s.repo.Create(spent)

	if err != nil {
		return domain.Spent{}, err
	}

	return *spent, nil
}

func (s *SpentService) GetSpents(userID uint) ([]domain.Spent, error) {
	return s.repo.GetSpents(userID)
}

func (s *SpentService) GetSpent(userID uint, spentID uint) (domain.Spent, error) {
	return s.repo.GetSpent(userID, spentID)
}

func (s *SpentService) DeleteSpent(userID uint, spentID uint) error {
	return s.repo.DeleteSpent(userID, spentID)
}

func (s *SpentService) UpdateSpent(spent domain.Spent) error {
	return s.repo.UpdateSpent(spent)
}