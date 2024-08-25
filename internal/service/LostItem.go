package service

import (
	"go-cleanarch/pkg/domain"
)

type LostItemService struct {
	lostItemRepository domain.LostItemRepository
}