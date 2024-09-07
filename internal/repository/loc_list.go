package repository

import (
	"errors"
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

type TempleLocList struct {
	TempleId     int
	TempleName   string
	Loc          string
	MainDeity    int
	History      string
	WorshipOrder string
	InCharge     string
	LinkRef      string
	NumsOfSubId  int
}

func (l *TempleLocList) TableName() string {
	return "m_m_list"
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

	return temple.TempleName, nil
}

func (r *postgresLocListRepository) GetNumOfSubLocByLocId(locId int) int {
	var templeList []TempleLocList
	r.db.Find(&templeList)
	for _, subLoc := range templeList {
		if subLoc.TempleId == locId {
			return subLoc.NumsOfSubId
		}
	}
	return -1
}

func (r *postgresLocListRepository) GetMainBadgeByLocationId(locationId int) (badge domain.Badge, err error) {
	var temple TempleLocList

	err = r.db.Where(&TempleLocList{TempleId: locationId}).Find(&temple).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	return domain.Badge{
		IconPath:    "", //TODO
		Description: "", //TODO
	}, nil
}

// get the quantity of sublocations in a location
func (r *postgresLocListRepository) GetSubLocQuantity(locationId int) (quantity int, err error) {
	//TODO
	var count int64
	err = r.db.Model(&TempleLocList{}).Where("temple_id = ?", 1).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil

}
