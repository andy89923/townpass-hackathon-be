package repository

import (
	"go-cleanarch/pkg/domain"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

type TbMap struct {
	LocationId int
	TbName     string
}

func (l *TbMap) TableName() string {
	return "m_m_list"
}

//------------------------------------------------

type postgresTbMapRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPostgresTbMApRepository(db *gorm.DB, logger *zap.Logger) domain.TbMapRepository {
	return &postgresTbMapRepository{
		db:     db,
		logger: logger,
	}
}

func (r *postgresTbMapRepository) GETXXX() error {
	//TODO
	return nil
}
