package domain


type MajorMinor uint32 // MajorMinor is a type for major and minor ，which are 4 bytes 


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
