package repository

import (
	"errors"
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

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
		Id : lostItemModel.ID,
		LostTime:    lostItemModel.LostTime,
		Kind:        lostItemModel.Kind,
		PropertyName: lostItemModel.PropertyName,
		Location:    lostItemModel.Location,
		PhoneNumber: lostItemModel.PhoneNumber,
	}, nil
}

func (r *postgresLostItemRepository) GetAll() ([]*domain.LostItem, error) {
	var lostItemsModel []*LostItem
	
	result := r.db.Find(&lostItemsModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	lostItems := make([]*domain.LostItem, 0, len(lostItemsModel))
	
	r.logger.Debug("GetAll: lostItems", zap.Int("len", len(lostItemsModel)))
	for _, lostItemModel := range lostItemsModel {
		lostItems = append(lostItems, &domain.LostItem{
			Id : lostItemModel.ID,
			LostTime:    lostItemModel.LostTime,
			Kind:        lostItemModel.Kind,
			PropertyName: lostItemModel.PropertyName,
			Location:    lostItemModel.Location,
			PhoneNumber: lostItemModel.PhoneNumber,
		})
		r.logger.Debug("GetAll: ", zap.Any("lostItems", lostItemModel))
	}

	
	return lostItems, nil
}

func (r *postgresLostItemRepository) GetByID(id uint) (*domain.LostItem, error) {
	var lostItemModel LostItem
	
	result := r.db.First(&lostItemModel, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	lostItem := &domain.LostItem{
		Id : lostItemModel.ID,
		LostTime:    lostItemModel.LostTime,
		Kind:        lostItemModel.Kind,
		PropertyName: lostItemModel.PropertyName,
		Location:    lostItemModel.Location,
		PhoneNumber: lostItemModel.PhoneNumber,
	}

	r.logger.Debug("GetByID: ", zap.Any("lostItem", lostItem))
	return lostItem, nil
}

func (r *postgresLostItemRepository) Update(lostItem *domain.LostItem) error {
	lostItemModel := LostItem{
		Model: gorm.Model{
			ID: lostItem.Id,
		},
		LostTime:  lostItem.LostTime,
		Kind:      lostItem.Kind,
		PropertyName: lostItem.PropertyName,
		Location:  lostItem.Location,
		PhoneNumber: lostItem.PhoneNumber,
	}

	result := r.db.Save(&lostItemModel)
	if result.Error != nil {
		return result.Error
	}

	r.logger.Debug("Update: ", zap.Any("lostItem", lostItemModel))
	lostItem.Id = lostItemModel.ID
	lostItem.LostTime = lostItemModel.LostTime
	lostItem.Kind = lostItemModel.Kind
	lostItem.PropertyName = lostItemModel.PropertyName
	lostItem.Location = lostItemModel.Location
	lostItem.PhoneNumber = lostItemModel.PhoneNumber

	return nil
}

func (r *postgresLostItemRepository) Delete(id uint) error {
	result := r.db.Delete(&LostItem{}, id)
	if result.Error != nil {
		return result.Error
	}

	r.logger.Debug("Delete: ", zap.Uint("id", id))

	return nil
}