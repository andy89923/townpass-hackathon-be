package repository

import (
	// "errors"
	"fmt"
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

type SubLocList struct {
	TempleId    int
	TempleName  string
	SubTempleId int
	Deity       string
	Description string
}

func (s *SubLocList)TableName() string{
	return "temple_subloc_list"
}

//--------------------------------------
type postgresSubLocListRepository struct {
	db *gorm.DB
	logger *zap.Logger
}


func NewPostgresSubLocListRepository(db *gorm.DB, logger *zap.Logger) domain.SubLocListRepository {
	return &postgresSubLocListRepository{
		db: db,
		logger: logger,
	}
}

// Get all sub locations (list) of a location
func (r *postgresSubLocListRepository) GetSubLoc(locId int) ([]domain.SubBadge, error) {
	var subBadges []domain.SubBadge
	var subTemples []SubLocList
	r.db.Where(&SubLocList{TempleId: locId}).Find(&subTemples)


	for _, sub := range(subTemples){
		subBadges = append(subBadges, domain.SubBadge{
			SubId: sub.SubTempleId,
			Badge: domain.Badge{
				IconPath: fmt.Sprint(locId) + fmt.Sprint(sub.SubTempleId), //TODO
				Description: "", //TODO
			},
		})
	}

	
	return subBadges, nil
}