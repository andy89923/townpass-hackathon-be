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

func NewPostfresqlVisitLogRepository(db *gorm.DB, logger *zap.Logger) domain.VisitLogRepository {
	return &postgresVisitLogRepository{
		db:     db,
		logger: logger,
	}
}

type postgresVisitLogRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (p *postgresVisitLogRepository) AddVisitLog(visitLog domain.VisitLog) (*domain.VisitLog, error ){
	//TODO
	visitLogModel := VisitLog{
		UserId:        visitLog.UserId,
		LocationId:    visitLog.LocationId,
		SubLocationId: visitLog.SubLocationId,
	}

	result := p.db.Create(&visitLogModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &domain.VisitLog{
		UserId:        visitLogModel.UserId,
		LocationId:    visitLogModel.LocationId,
		SubLocationId: visitLogModel.SubLocationId,
	}, nil
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
