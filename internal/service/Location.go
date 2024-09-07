package service

import (
	"fmt"
	"go-cleanarch/pkg/domain"

	"go.uber.org/zap"
)

type LocationService struct {
	locationRepository   domain.LocationRepository
	locListRepository    domain.LocListRepository
	subLocListRepository domain.SubLocListRepository
	visitLogRepository   domain.VisitLogRepository
	tbMapRepository      domain.TbMapRepository
	logger               *zap.Logger
}

func NewBadgeService(locationRepository domain.LocationRepository,
	locListRepository domain.LocListRepository,
	subLocListRepository domain.SubLocListRepository,
	visitLogRepository domain.VisitLogRepository,
	tbMapRepository domain.TbMapRepository,
	logger *zap.Logger) *LocationService {
	return &LocationService{
		locationRepository:   locationRepository,
		locListRepository:    locListRepository,
		subLocListRepository: subLocListRepository,
		visitLogRepository:   visitLogRepository,
		tbMapRepository:      tbMapRepository,
		logger:               logger,
	}
}

func (s *LocationService) GetBadge(mm domain.MajorMinor, id int) (*domain.Location, error) {
	s.logger.Debug("[Service] GetBadge")

	resp := domain.Location{}
	resp.MajorMinor = domain.MajorMinor(mm)

	// use MM to get locationId, sublocationId
	locationId, sublocationId, err := s.locationRepository.GetLocationByMM(mm)
	if err != nil {
		s.logger.Debug("[Service] GetBadge GetLocationByMM error")
		return nil, fmt.Errorf("[Service] GetBadge GetLocationByMM error: %v", err)
	}

	// GetTableByLocationId from tbMap to ensure which table to use

	resp.LocationName, err = s.locListRepository.GetNameByLocation(locationId)
	if err != nil {
		s.logger.Debug("[Service] GetBadge GetNameByLocation error")
		return nil, fmt.Errorf("[Service] GetBadge GetNameByLocation error: %v", err)
	}

	// record current sublocation to visit_log
	err = s.visitLogRepository.AddVisitLog(mm, id, locationId, sublocationId)
	if err != nil {
		s.logger.Debug("[Service] GetBadge AddVisitLog error")
		return nil, fmt.Errorf("[Service] GetBadge AddVisitLog error: %v", err)
	}

	// get all sublocation info by locationId
	subBadgesFromDB, err := s.subLocListRepository.getSubLocListByLocId(locationId)
	if err != nil {
		s.logger.Debug("[Service] GetBadge getSubLocListByLocId error")
		return nil, fmt.Errorf("[Service] GetBadge getSubLocListByLocId error: %v", err)
	}
	// get user's visited record
	visitList, err := s.visitLogRepository.GetVisitedSubLocIdsByUserLocInfo(id, locationId)
	if err != nil {
		s.logger.Debug("[Service] GetBadge GetVisitedSubLocIdsByUserLocInfo error")
		return nil, fmt.Errorf("[Service] GetBadge GetVisitedSubLocIdsByUserLocInfo error: %v", err)
	}
	// compare sublocation info with user's visited record
	visitedMap := make(map[int]bool)
	for _, visit := range visitList {
		visitedMap[visit] = true
	}

	var respSubBadges []domain.SubBadge
	countProgress := 0
	for _, subBadge := range subBadgesFromDB {
		badge := domain.Badge{
			IconPath:    subBadge.IconPath,
			Description: subBadge.Description,
		}

		if visitedMap[subBadge.SubId] {
			badge.Aquired = true
			countProgress++
		} else {
			badge.Aquired = false
		}

		subBadgeTmp := domain.SubBadge{
			Badge: badge,
			SubId: subBadge.SubId,
		}

		respSubBadges = append(respSubBadges, subBadgeTmp)
	}

	resp.SubBadge = respSubBadges

	resp.Progress = countProgress
	resp.TotalProgress, err = s.locationRepository.GetNumOfSubLocByLocId(locationId)
	if err != nil {
		s.logger.Debug("[Service] GetBadge GetNumOfSubLocByLocId error")
		return nil, fmt.Errorf("[Service] GetBadge GetNumOfSubLocByLocId error: %v", err)
	}

	// get main badge info by locationId
	mainBadgeDB, err := s.locationRepository.GetMainBadgeByLocationId(locationId)
	if err != nil {
		s.logger.Debug("[Service] GetBadge GetMainBadgeByLocationId error")
		return nil, fmt.Errorf("[Service] GetBadge GetMainBadgeByLocationId error: %v", err)
	}

	mainBadge := domain.Badge{
		IconPath:    mainBadgeDB.IconPath,
		Description: mainBadgeDB.Description,
	}
	if resp.Progress == resp.TotalProgress {
		mainBadge.Aquired = true
	} else {
		mainBadge.Aquired = false
	}
	resp.MainBadge = mainBadge

	return &resp, nil
}
