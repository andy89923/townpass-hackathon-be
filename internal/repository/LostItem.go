package repository

import (

	"gorm.io/gorm"
	"go-cleanarch/pkg/domain"

	"go.uber.org/zap"
)

type LostItem struct {
	gorm.Model

	LostTime  string
	Kind      string
	PropertyName string
	Location  string
	PhoneNumber string
}


type postgresLostItemRepository struct {
	db *gorm.DB
	logger *zap.Logger
}

func NewPostgresLostItemRepository(db *gorm.DB, logger *zap.Logger) domain.LostItemRepository {
	return &postgresLostItemRepository{
		db: db,
		logger: logger,
	}
}

func (r *postgresLostItemRepository) Create(lostItem *domain.LostItem) (*domain.LostItem, error) {
	lostItemModel := LostItem{ // convert domain.LostItem to repository.LostItem
		LostTime:  lostItem.LostTime,
		Kind:      lostItem.Kind,
		PropertyName: lostItem.PropertyName,
		Location:  lostItem.Location,
		PhoneNumber: lostItem.PhoneNumber,
	}

	result := r.db.Create(&lostItemModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &domain.LostItem{
		LostTime:    lostItemModel.LostTime,
		Kind:        lostItemModel.Kind,
		PropertyName: lostItemModel.PropertyName,
		Location:    lostItemModel.Location,
		PhoneNumber: lostItemModel.PhoneNumber,
	}, nil
}
