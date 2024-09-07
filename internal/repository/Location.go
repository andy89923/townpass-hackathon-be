package repository

import (
	"errors"
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

	"go.uber.org/zap"
)


type LocationTable struct {
	gorm.Model

	mm domain.MajorMinor
	location uint8
	subLocation uint8
}

type postgresBadgeRepository struct {
	db *gorm.DB
	logger *zap.Logger
}

func NewPostgresBadgeRepository(db *gorm.DB, logger *zap.Logger) domain.LostItemRepository {
	return &postgresLostItemRepository{
		db: db,
		logger: logger,
	}
}

func (r *postgresBadgeRepository) QueryLocationByMM(mm domain.MajorMinor) (string, error) {
	// var location LocationTable
	// err := r.db.Where("mm = ?", mm).First(&location).Error
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return "", nil // record not found
	// 	}
	// 	return "", err // other error occurred
	// }
	// return location.Name, nil
}
