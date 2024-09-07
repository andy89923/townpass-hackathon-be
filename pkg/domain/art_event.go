package domain

type ArtSubEvent struct {
	SubeventId  int    `json:"subeventId" form:"subevent_id"`
	EventId     int    `json:"eventId" form:"event_id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

type ArtEvent struct {
	EventId     int           `json:"eventId" form:"event_id"`
	EventName   string        `json:"name" form:"name"`
	Description string        `json:"description" form:"description"`
	PlaceId     int           `json:"placeId" form:"place_id"`
	Subevents   []ArtSubEvent `json:"subevents" form:"subevents"`
}

type ArtEventRepository interface {
	GetEventByMM(mm MajorMinor) (event ArtEvent, err error)
}
