package service

import (
	"go-cleanarch/pkg/domain"
)

type LostItemService struct {
	lostItemRepository domain.LostItemRepository
}

func NewLostItemService(lostItemRepository domain.LostItemRepository) *LostItemService {
	return &LostItemService{lostItemRepository: lostItemRepository}
}

func (s *LostItemService) AddNewLostItem(lostItem *domain.LostItem) (*domain.LostItem, error) {
	lostItem, err := s.lostItemRepository.Create(lostItem)

	if err != nil {
		return nil, err
	}

	return lostItem, nil
}