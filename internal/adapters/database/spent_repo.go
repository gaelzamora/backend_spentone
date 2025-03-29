package database

import (
	"github.com/gaelzamora/spent-one/internal/domain"
	"github.com/gaelzamora/spent-one/internal/ports"
	"gorm.io/gorm"
)

type SpentRepositoryImpl struct {
	DB *gorm.DB
}

func NewSpentRepositoryImpl(db *gorm.DB) ports.SpentRepository {
	return &SpentRepositoryImpl{DB: db}
}

func (s SpentRepositoryImpl) Create(spent *domain.Spent) error {
	err := s.DB.Create(&spent).Error
	if err != nil {
		return err
	}

	return nil
}

func (s SpentRepositoryImpl) GetSpents(userID uint) ([]domain.Spent, error) {
	var spents []domain.Spent
	err := s.DB.Where("user_id = ?", userID).Find(&spents).Error

	return spents, err
}

func (s SpentRepositoryImpl) GetSpent(userID uint, spentID uint) (domain.Spent, error) {
	var spent domain.Spent
	err := s.DB.Where("user_id = ? AND id = ?", userID, spentID).First(&spent).Error

	return spent, err
}

func (s SpentRepositoryImpl) DeleteSpent(userID uint, spentID uint) error {
	result := s.DB.Where("user_id = ? AND id = ?", userID, spentID).Delete(&domain.Spent{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (s SpentRepositoryImpl) UpdateSpent(spent domain.Spent) error {
	return s.DB.Save(&spent).Error
}
