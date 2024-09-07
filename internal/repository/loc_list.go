package repository

import (
	"errors"
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

type LocList struct {
	templeId 		int 	`gorm:"column:temple_id"`
	templeName 		string  `gorm:"column:temple_name"`
	loc				string	`gorm:"column:loc"`
	mainDeity 		int		`gorm:"column:main_deity"`
	history 		string	`gorm:"column:history"`
	worshipOrder 	string	`gorm:"column:worship_order"`
	inCharge 		string	`gorm:"column:in_charge"`
	linkRef 		string	`gorm:"column:link_ref"`
	numsOfSubId 	int		`gorm:"column:nums_of_sub_id"`
}

type postgresLocListRepository struct {
	db *gorm.DB
	logger *zap.Logger
}


func NewPostgresLocListRepository(db *gorm.DB, logger *zap.Logger) domain.LocListRepository {
	return &postgresLocListRepository{
		db: db,
		logger: logger,
	}
}

func (r *postgresLocationRepository) GetNameByLocation(locationId int) (name string, err error) {
	var temple LocList
	
	err = r.db.Where(&LocList{templeId: locationId}).Find(&temple).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", domain.ErrNotFound	
	}

	return temple.templeName, nil
}

func (r *postgresLostItemRepository) GetNumOfSubLocByLocId(locId int) int {
	var templeList []LocList
	r.db.Find(&templeList)
	for _, subLoc := range templeList {
		if subLoc.templeId == locId {
			return subLoc.numsOfSubId
		}
	}
	return -1
}

func (r *postgresLocationRepository) GetMainBadgeByLocationId(locationId int) (badge domain.Badge, err error) {
	var temple LocList
	
	err = r.db.Where(&LocList{templeId: locationId}).Find(&temple).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 
	}


	return domain.Badge{
		IconPath: "", //TODO
		Description: "", //TODO
	}, nil
}

// get the quantity of sublocations in a location
func (r *postgresLocationRepository) GetSubLocQuantity(locationId int) (quantity int, err error) {
	//TODO
	var count int64
	err = r.db.Model(&LocList{}).Where("temple_id = ?", 1).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil

}