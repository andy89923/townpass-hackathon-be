package service

import (
	"fmt"
	"strconv"
	"go-cleanarch/pkg/domain"
)

type CollectionService struct {
	locListRepository    domain.LocListRepository
	visitLogRepository   domain.VisitLogRepository
}

func NewCollectionService(locListRepository domain.LocListRepository,
	visitLogRepository domain.VisitLogRepository) *CollectionService {
	return &CollectionService{
		locListRepository:  locListRepository,
		visitLogRepository: visitLogRepository,
	}
}

func (s *CollectionService) GetCollections(userId int) ([]domain.Collection, error) {
	var collections []domain.Collection
	visitedLocations, err := s.visitLogRepository.GetVisitedLocIdsByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("[Service] GetCollections GetVisitedLocIdsByUserId error: %v", err)
	}

	for _, locId := range visitedLocations {
		locName, err := s.locListRepository.GetNameByLocation(locId)
		if err != nil {
			return nil, fmt.Errorf("[Service] GetCollections GetLocName error: %v", err)
		}

		collections = append(collections, domain.Collection{
			IconPath:    strconv.Itoa(locId) + "_0",
			LocationName: locName,
		})
	}

	return collections, nil
}