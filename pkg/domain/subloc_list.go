package domain

// "go-cleanarch/pkg/domain"
// "gorm.io/gorm"
// "go.uber.org/zap"

type SubLocListRepository interface {
	GetSubLocListByLocId(locId int) ([]SubBadge, error)
}
