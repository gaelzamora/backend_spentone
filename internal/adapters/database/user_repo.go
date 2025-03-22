package database

import (
	"github.com/gaelzamora/spent-one/internal/domain"
	"github.com/gaelzamora/spent-one/internal/ports"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	result := r.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryImpl) FindByID(id uint) (*domain.User, error) {
	var user domain.User

	result := r.DB.Select("id", "username", "created_at").First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
