package repository

import (
	// "errors"
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

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