package domain

type MajorMinor uint32

type Location struct {
	MajorMinor MajorMinor // MajorMinor is a type for major and minor ，which are 4 bytes
	UserId     int        `json:"userId" form:"user_id"`
	// LocationName string     `json:"name" form:"name"`
	Name        string `json:"name" form:"name"`
	Progress    int    `json:"progress" form:"progress"`
	NumsOfSubId int    `json:"total" form:"total"`

	MainBadge Badge      `json:"mainBadge" form:"main_badge"`
	SubBadge  []SubBadge `json:"subBadges" form:"sub_badge"`
}

type Badge struct {
	IconPath    string            `json:"iconPath" form:"icon_path"`
	Aquired     bool              `json:"aquired" form:"aquired"`
	Description TempleDescription `json:"description" form:"description"`
}

type SubBadge struct {
	Badge
	SubId int `json:"sub_id" form:"sub_id"`
	SubDescription SubTempleDescription
}

type TempleDescription struct {
	MainDeity    string
	History      string
	WorshipOrder string
	InCharge     string
	LinkRef      string
}

type SubTempleDescription struct {
	Deity       string
	Description string
}

type LocationRepository interface {
	//TODO
	GetLocationByMM(mm MajorMinor) (locationId int, subLocationId int, err error)
	Create(location *Location, locationId int, subLocationId int) error
}
