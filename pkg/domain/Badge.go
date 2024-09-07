package domain

type MajorMinor struct {
	Major int
	Minor int
}

type Location_DB struct {
	Location    int
	name        string
	Description string
}

type LocationSublocation_DB struct {
	Location    int
	Sublocation int
	Description string
}

type VisitLog_DB struct {
	UserId      int
	Location    int
	Sublocation int
}
