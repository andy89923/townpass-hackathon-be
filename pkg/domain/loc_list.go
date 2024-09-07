package domain



type LocListRepository interface {
	// TODO
	GetNameByLocation(locationId int) (name string, err error) 
	GetNumOfSubLocByLocId(locId int) int
	GetMainBadgeByLocationId(locationId int) (Badge, error)
	GetSubLocQuantity(locationId int) (quantity int, err error) 
}


