package domain

type VisitLog struct {
	UserId        int
	LocationId    int
	SubLocationId int
}

type VisitLogRepository interface {
	AddVisitLog(mm MajorMinor, id int, locationId int, sublocationId int) error
	GetVisitedSubLocIdsByUserLocInfo(userId int, locationId int) []int
}
