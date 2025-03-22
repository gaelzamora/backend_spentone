package database

import "gorm.io/gorm"

type SpentRepositoryImpl struct {
	DB *gorm.DB
}

func NewSpentRepositoryImpl(db *gorm.DB) SpentRepositoryImpl {
	return &SpentRepositoryImpl{DB: db}
}

