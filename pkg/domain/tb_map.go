package domain

// "go-cleanarch/pkg/domain"
// "gorm.io/gorm"
// "go.uber.org/zap"

type TbMapRepository interface {
	GetTableByLocationId(LocationId int) (string, error)
}
