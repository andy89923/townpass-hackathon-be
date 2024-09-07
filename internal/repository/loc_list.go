package repository

import (
	"errors"
	"fmt"
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

type TempleLocList struct {
	TempleId     int
	TempleName   string
	Loc          string
	MainDeity    string
	History      string
	WorshipOrder string
	InCharge     string
	LinkRef      string
	NumsOfSubId  int
}

func (l *TempleLocList) TableName() string {
	return "temple_loc_list"
}

//------------------------------------------------

type postgresLocListRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPostgresLocListRepository(db *gorm.DB, logger *zap.Logger) domain.LocListRepository {
	return &postgresLocListRepository{
		db:     db,
		logger: logger,
	}
}

func (r *postgresLocListRepository) GetNameByLocation(locationId int) (name string, err error) {
	var temple TempleLocList

	err = r.db.Where(&TempleLocList{TempleId: locationId}).Find(&temple).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", domain.ErrNotFound
	}
	fmt.Println("temple.TempleName: ", temple.TempleName)

	return temple.TempleName, nil
}

func (r *postgresLocListRepository) GetMainBadgeByLocationId(locationId int) (badge domain.Badge, err error) {
	var temple TempleLocList

	err = r.db.Where(&TempleLocList{TempleId: locationId}).Find(&temple).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	return domain.Badge{
		IconPath:    fmt.Sprint(locationId) + "_0", 
		Description: domain.TempleDescription{
			MainDeity:    temple.MainDeity,
			History:      temple.History,
			WorshipOrder: temple.WorshipOrder,
			InCharge:     temple.InCharge,
			LinkRef:      temple.LinkRef,
		},
	}, nil
}

// get the quantity of sublocations in a location
func (r *postgresLocListRepository) GetSubLocQuantity(locationId int) (quantity int, err error) {
	//TODO
	var temple TempleLocList
	err = r.db.Where("temple_id = ?", locationId).First(&temple).Error
	if err != nil {
		return 0, err
	}
	return temple.NumsOfSubId, nil
}
