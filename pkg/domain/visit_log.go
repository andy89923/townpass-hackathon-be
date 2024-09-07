package domain

type VisitLog struct {
	UserId   int
	LocId    int
	SubLocId int
}

type VisitLogRepository interface {
	AddVisitLog(visitLog VisitLog) (*VisitLog, error )
	GetVisitedSubLocIdsByUserLocInfo(userId int, locationId int) ([]int, error)
}
