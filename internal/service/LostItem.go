package service

import (
	"go-cleanarch/pkg/domain"

	"go.uber.org/zap"
)

type LostItemService struct {
	lostItemRepository domain.LostItemRepository
	logger *zap.Logger
}

func NewLostItemService(lostItemRepository domain.LostItemRepository, logger *zap.Logger) *LostItemService {
	return &LostItemService{
		lostItemRepository: lostItemRepository,
		logger: logger,
	}
}

func (s *LostItemService) AddNewLostItem(lostItem *domain.LostItem) (*domain.LostItem, error) {
	s.logger.Debug("AddNewLostItem")

	lostItem, err := s.lostItemRepository.Create(lostItem)

	if err != nil {
		return nil, err
	}

	return lostItem, nil
}