package repository

import (
	// "errors"
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

type SubLocList struct {
	TempleId    int
	TempleName  string
	SubTempleId string
	Deity       string
	Description string
}

func (s *SubLocList)TableName() string{
	return "temple_subloc_list"
}

//--------------------------------------
type postgresSubLocListRepository struct {
	db *gorm.DB
	logger *zap.Logger
}


func NewPostgresSubLocListRepository(db *gorm.DB, logger *zap.Logger) domain.SubLocListRepository {
	return &postgresSubLocListRepository{
		db: db,
		logger: logger,
	}
}