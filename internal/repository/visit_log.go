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

func NewPostgresVisitLogRepository(db *gorm.DB, logger *zap.Logger) domain.VisitLogRepository {
	return &postgresVisitLogRepository{
		db:     db,
		logger: logger,
	}
}

type postgresVisitLogRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPostgresVisitLogRepository(db *gorm.DB, logger *zap.Logger) domain.VisitLogRepository {
	return &postgresVisitLogRepository{
		db:     db,
		logger: logger,
	}
}

func (p *postgresVisitLogRepository) AddVisitLog(visitLog VisitLog) error {
func (p *postgresVisitLogRepository) AddVisitLog(visitLog domain.VisitLog) (*domain.VisitLog, error ){
	//TODO
	visitLogModel := VisitLog{
		UserId:        visitLog.UserId,
		LocId:    visitLog.LocationId,
		SubLocId: visitLog.SubLocationId,
	}

	result := p.db.Create(&visitLogModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &domain.VisitLog{
		UserId:        visitLogModel.UserId,
		LocationId:    visitLogModel.LocId,
		SubLocationId: visitLogModel.SubLocId,
	}, nil
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
