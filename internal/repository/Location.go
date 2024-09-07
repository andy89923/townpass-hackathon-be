package repository

import (
	// "errors"
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

type postgresLocationRepository struct {
	db *gorm.DB
	logger *zap.Logger
}

func NewPostgresBadgeRepository(db *gorm.DB, logger *zap.Logger) domain.LocationRepository {
	return &postgresLocationRepository{
		db: db,
		logger: logger,
	}
}

func (r *postgresLocationRepository) GetLocationByMM(mm domain.MajorMinor) (locationId int, subLocationId int, err error) {
	//TODO
	return 0,0, nil
}

func (r *postgresLocationRepository) GetNameByLocation(locationId int) (name string, err error) {
	//TODO
	
	return "", nil
}

func (r *postgresLocationRepository) GetMainBadgeByLocationId(locationId int) (badge domain.Badge, err error) {
	//TODO
	return domain.Badge{}, nil
}

func (r *postgresLocationRepository) GetSubLocQuantity(locationId int) (quantity int, err error) {
	//TODO
	return 0, nil
}