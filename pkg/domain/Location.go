package domain

type MajorMinor uint32

type Location struct {
	MajorMinor MajorMinor // MajorMinor is a type for major and minor ，which are 4 bytes
	UserId     int        `json:"user_id" form:"user_id"`
	// LocationName string     `json:"name" form:"name"`
	Name        string `json:"name" form:"name"`
	Progress    int    `json:"progress" form:"progress"`
	NumsOfSubId int    `json:"total" form:"total"`

	MainBadge Badge      `json:"main_badge" form:"main_badge"`
	SubBadge  []SubBadge `json:"sub_badges" form:"sub_badge"`
}

type Badge struct {
	IconPath    string `json:"icon_path" form:"icon_path"`
	Aquired     bool   `json:"aquired" form:"aquired"`
	Description string `json:"description" form:"description"`
}

type SubBadge struct {
	Badge
	SubId int `json:"sub_id" form:"sub_id"`
}

type LocationRepository interface {
	//TODO
	GetLocationByMM(mm MajorMinor) (locationId int, subLocationId int, err error)
	Create(location *Location, locationId int, subLocationId int) error
}
