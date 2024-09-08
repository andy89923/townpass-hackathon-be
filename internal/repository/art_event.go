package repository

import (
	"encoding/json"
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
	PlaceId     int
}

func (l *ArtEvent) TableName() string {
	return "art_event_list"
}

type ArtSubEvent struct {
	SubeventId  int `gorm:"primaryKey"`
	EventId     int `gorm:"foreignKey:EventId"`
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
	eventModel := ArtEvent{}
	result = r.db.Where(&ArtEvent{EventId: subevent.EventId}).First(&eventModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}

	jsonTemp, err := json.Marshal(eventModel)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonTemp, event)
	if err != nil {
		return nil, err
	}

	event.EventName = eventModel.EventName

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

	fmt.Printf("event id: %d\n", subevent.EventId)
	fmt.Printf("event: %+v\n", event)

	return event, nil
}
