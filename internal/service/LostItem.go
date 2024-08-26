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

func (s *LostItemService) GetAllLostItems() ([]*domain.LostItem, error) {
	lostItems, err := s.lostItemRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return lostItems, nil
}

func (s *LostItemService) GetLostItemById(id uint) (*domain.LostItem, error) {
	lostItem, err := s.lostItemRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return lostItem, nil
}

func (s *LostItemService) UpdateLostItem(lostItem *domain.LostItem) (error) {
	err := s.lostItemRepository.Update(lostItem)
	if err != nil {
		return err
	}

	return nil
}

func (s *LostItemService) DeleteLostItem(id uint) error {
	err := s.lostItemRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}