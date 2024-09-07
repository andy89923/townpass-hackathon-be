package repository

import (
	"errors"
	"fmt"
	"go-cleanarch/pkg/domain"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ArtEvent struct {
	EventId     int `gorm:"primaryKey"`
	EventName   string
	Description string
	placeId     int
}

func (l *ArtEvent) TableName() string {
	return "art_event_list"
}

type ArtSubEvent struct {
	SubeventId  int `gorm:"primaryKey"`
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

func (r *postgresArtEventListRepository) GetEventBySubeventId(id int) (event *domain.ArtEvent, err error) {
	var subevent ArtSubEvent
	result := r.db.Where(&ArtSubEvent{SubeventId: id}).First(&subevent)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}

	event = new(domain.ArtEvent)
	result = r.db.Where(&ArtEvent{EventId: subevent.EventId}).First(event)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}

	var subevents []ArtSubEvent
	result = r.db.Where(&ArtSubEvent{EventId: subevent.EventId}).Find(&subevents)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}

	for _, subevent := range subevents {
		event.Subevents = append(event.Subevents, domain.ArtSubEvent{
			SubeventId:  subevent.SubeventId,
			EventId:     subevent.EventId,
			Name:        subevent.Name,
			Description: subevent.Description,
		})
	}

	fmt.Printf("event: %+v\n", event)

	return event, nil
}
