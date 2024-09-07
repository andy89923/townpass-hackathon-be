package repository

import (
	// "errors"
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


func NewPostgresLocListRepository(db *gorm.DB, logger *zap.Logger) domain.LocationRepository {
	return &postgresLocListRepository{
		db: db,
		logger: logger,
	}
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