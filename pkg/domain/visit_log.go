package domain

type VisitLogRepository interface {
	AddVisitLog(visitLog VisitLog) (*VisitLog, error )
	GetVisitedSubLocIdsByUserLocInfo(userId int, locationId int) ([]int, error)
}
