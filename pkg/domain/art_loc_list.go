package domain

type ArtLocListRepository interface {
	GetLocationByPlaceId(placeId int) (location Location, err error)
}
