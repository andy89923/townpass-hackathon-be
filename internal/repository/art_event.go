package repository

import (
	"go-cleanarch/pkg/domain"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ArtEvent struct {
	EventId     int
	EventName   string
	Description string
	placeId     int
}

func (l *ArtEvent) TableName() string {
	return "art_event_list"
}

type ArtSubEvent struct {
	SubeventId  int
	EventId     int
	Name        string
	Description string
}

func (l *ArtSubEvent) TableName() string {
	return "art_subevent_list"
}

type postgresArtEventListRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPostgresArtEventListRepository(db *gorm.DB, logger *zap.Logger) domain.ArtEventRepository {
	return &postgresArtEventListRepository{
		db:     db,
		logger: logger,
	}
}

func (r *postgresArtEventListRepository) GetEventByMM(mm domain.MajorMinor) (event domain.ArtEvent, err error) {
	// TODO: get event and all the related subevents
}
