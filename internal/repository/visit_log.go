package repository

import (
	// "errors"
	"errors"
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

type VisitLog struct {
	gorm.Model

	UserId   int
	LocId    int
	SubLocId int
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

func (r *postgresVisitLogRepository) GetVisitedSubLocIdsByUserLocInfo(userId int, locationId int) (visitedList []int, err error) {
	var visitLogList []VisitLog
	var visitedSubLocIds []int
	result := r.db.Where("user_id = ? AND loc_id = ?", userId, locationId).Find(&visitLogList)
	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	for _, visitLog := range visitLogList {
		visitedSubLocIds = append(visitedSubLocIds, visitLog.SubLocId)
	}
	return visitedSubLocIds, nil
}
