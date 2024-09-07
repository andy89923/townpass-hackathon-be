package repository

import (
	// "errors"
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

type SubLocList struct {
	templeId    int    `gorm:"column:temple_id"`
	templeName  string `gorm:"column:temple_name"`
	subTempleId string `gorm:"column:sub_temple_id"`
	deity       string `gorm:"column:deity"`
	description string `gorm:"column:description"`
}

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