package repository

import (
	// "errors"
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

type VisitLog struct {
	gorm.Model

	UserId        int
	LocationId    int
	SubLocationId int
}

type postgresVisitLogRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (p *postgresVisitLogRepository) AddVisitLog(visitLog domain.VisitLog) error {
	//TODO
	return nil
}

func (p *postgresVisitLogRepository) IsEventExist(userId int, locationId int) (bool, error) {
	//TODO
	return false, nil
}

func (r *postgresVisitLogRepository) GetVisitedSubLocIdsByUserLocInfo(userId int, locationId int) []int {
	var visitLogList []VisitLog
	var visitedSubLocIds []int
	r.db.Find(&visitLogList)
	for _, visitLog := range visitLogList {
		if visitLog.UserId == userId && visitLog.LocationId == locationId {
			visitedSubLocIds = append(visitedSubLocIds, visitLog.SubLocationId)
		}
	}
	return visitedSubLocIds
}
