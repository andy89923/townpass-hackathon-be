package repository

import (
	"gorm.io/gorm"
	"go-cleanarch/pkg/domain"
)



type postgresLostItemRepository struct {
	db *gorm.DB
}

func NewPostgresLostItemRepository(db *gorm.DB) domain.LostItemRepository {
	return &postgresLostItemRepository{db: db}
}

func (r *postgresLostItemRepository) Create(lostItem *domain.LostItem) (*domain.LostItem, error) {
	// Todo: Implement this
	return nil, nil
}
