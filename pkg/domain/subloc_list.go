package domain

// "go-cleanarch/pkg/domain"
// "gorm.io/gorm"
// "go.uber.org/zap"

type SubLocList struct {
	templeId    int    `gorm:"column:temple_id"`
	templeName  string `gorm:"column:temple_name"`
	subTempleId string `gorm:"column:sub_temple_id"`
	deity       string `gorm:"column:deity"`
	description string `gorm:"column:description"`
}

type SubLocListRepository interface {
	// TODO
}
