package domain

type VisitLog struct {
	UserId        int
	LocationId    int
	SubLocationId int
}

type VisitLogRepository interface {
	AddVisitLog(visitLog VisitLog) (*VisitLog, error )
	GetVisitedSubLocIdsByUserLocInfo(userId int, locationId int) ([]int, error)
}
