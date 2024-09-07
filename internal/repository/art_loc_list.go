package repository

import (
	"errors"
	"go-cleanarch/pkg/domain"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ArtLocList struct {
	PlaceId     int `gorm:"primaryKey"`
	PlaceName   string
	Description string
	Loc         string
}

func (l *ArtLocList) TableName() string {
	return "art_loc_list"
}

type postgresArtLocListRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPostgresArtLocListRepository(db *gorm.DB, logger *zap.Logger) domain.ArtLocListRepository {
	return &postgresArtLocListRepository{
		db:     db,
		logger: logger,
	}
}

func (r *postgresArtLocListRepository) GetLocationByPlaceId(placeId int) (location *domain.Location, err error) {
	var artLocList ArtLocList

	result := r.db.Where(&ArtLocList{PlaceId: placeId}).First(&artLocList)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	location = new(domain.Location)
	location.Name = artLocList.PlaceName
	location.MainBadge = new(domain.Badge)
	location.MainBadge.Description.History = artLocList.Description

	return location, nil
}
