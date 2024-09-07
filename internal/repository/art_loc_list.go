package repository

import (
	"go-cleanarch/pkg/domain"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ArtLocList struct {
	PlaceId     int
	PlaceName   string
	Description string
	Loc         string
}

func (l *TempleLocList) TableName() string {
	return "art_loc_list"
}

type postgresArtLocListRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPostgresArtLocListRepository(db *gorm.DB, logger *zap.Logger) domain.LocListRepository {
	return &postgresLocListRepository{
		db:     db,
		logger: logger,
	}
}

func (r *postgresArtLocListRepository) GetLocationByPlaceId(placeId int) (location domain.Location, err error) {
	// TODO: get location
}
