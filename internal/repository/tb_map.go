package repository

import (
	"errors"
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

type TbMap struct {
	gorm.Model
	
	LocationId int
	TbName     string
}

func (l *TbMap) TableName() string {
	return "tbmap"
}

//------------------------------------------------

type postgresTbMapRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPostgresTbMapRepository(db *gorm.DB, logger *zap.Logger) domain.TbMapRepository {
	return &postgresTbMapRepository{
		db:     db,
		logger: logger,
	}
}

func (r *postgresTbMapRepository) GetTableByLocationId(locationId int) (string, error) {
	var tbMap TbMap
	result := r.db.Where(&TbMap{LocationId: locationId}).Find(&tbMap)
	err := result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", domain.ErrNotFound
	}
	return tbMap.TbName, nil
}
