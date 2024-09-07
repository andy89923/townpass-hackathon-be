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
	artLocListRepository domain.ArtLocListRepository
	artEventRepository   domain.ArtEventRepository
	logger               *zap.Logger
}

func NewBadgeService(locationRepository domain.LocationRepository,
	locListRepository domain.LocListRepository,
	subLocListRepository domain.SubLocListRepository,
	visitLogRepository domain.VisitLogRepository,
	tbMapRepository domain.TbMapRepository,
	artLocListRepository domain.ArtLocListRepository,
	artEventRepository domain.ArtEventRepository,
	logger *zap.Logger) *LocationService {
	return &LocationService{
		locationRepository:   locationRepository,
		locListRepository:    locListRepository,
		subLocListRepository: subLocListRepository,
		visitLogRepository:   visitLogRepository,
		tbMapRepository:      tbMapRepository,
		artLocListRepository: artLocListRepository,
		logger:               logger,
	}
}

const (
	TBMAP_TEMPLE = "temple"
	TBMAP_ART    = "art"
	TBMAP_SPORT  = "sport"
)

func (s *LocationService) GetBadge(mm domain.MajorMinor, id int) (*domain.Location, error) {
	s.logger.Debug("[Service] GetBadge")

	resp := domain.Location{}
	resp.MajorMinor = domain.MajorMinor(mm)
	resp.UserId = id

	// use MM to get locationId, sublocationId
	locationId, sublocationId, err := s.locationRepository.GetLocationByMM(mm)
	if err != nil {
		s.logger.Debug("[Service] GetBadge GetLocationByMM error")
		return nil, fmt.Errorf("[Service] GetBadge GetLocationByMM error: %v", err)
	}

	// GetTableByLocationId from tbMap to ensure which table to use
	tableName, err := s.tbMapRepository.GetTableByLocationId(locationId)
	if err != nil {
		s.logger.Debug("[Service] Get Table Name By Location ID error")
		return nil, fmt.Errorf("[Service] Get Table Name By Location ID error: %v", err)
	}

	if tableName == "temple" {
		resp.Name, err = s.locListRepository.GetNameByLocation(locationId)
		fmt.Println("resp.LocationName: ", resp.Name)
		if err != nil {
			s.logger.Debug("[Service] GetBadge GetNameByLocation error")
			return nil, fmt.Errorf("[Service] GetBadge GetNameByLocation error: %v", err)
		}
		visit_log := domain.VisitLog{
			UserId:   id,
			LocId:    locationId,
			SubLocId: sublocationId,
		}
		// record current sublocation to visit_log
		_, err = s.visitLogRepository.AddVisitLog(visit_log)
		if err != nil {
			s.logger.Debug("[Service] GetBadge AddVisitLog error")
			return nil, fmt.Errorf("[Service] GetBadge AddVisitLog error: %v", err)
		}

		// get all sublocation info by locationId
		subBadgesFromDB, err := s.subLocListRepository.GetSubLocListByLocId(locationId)
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
		resp.NumsOfSubId, err = s.locListRepository.GetSubLocQuantity(locationId)
		fmt.Print("resp.TotalProgress: ", resp.NumsOfSubId)
		if err != nil {
			s.logger.Debug("[Service] GetBadge GetNumOfSubLocByLocId error")
			return nil, fmt.Errorf("[Service] GetBadge GetNumOfSubLocByLocId error: %v", err)
		}

		// get main badge info by locationId
		mainBadgeDB, err := s.locListRepository.GetMainBadgeByLocationId(locationId)
		if err != nil {
			s.logger.Debug("[Service] GetBadge GetMainBadgeByLocationId error")
			return nil, fmt.Errorf("[Service] GetBadge GetMainBadgeByLocationId error: %v", err)
		}

		mainBadge := domain.Badge{
			IconPath:    mainBadgeDB.IconPath,
			Description: mainBadgeDB.Description,
		}
		if resp.Progress == resp.NumsOfSubId {
			mainBadge.Aquired = true
		} else {
			mainBadge.Aquired = false
		}
		resp.MainBadge = mainBadge
	} else if tableName == TBMAP_ART {
		event, err := s.artEventRepository.GetEventByMM(mm)
		if err != nil {
			s.logger.Debug("[Service] GetBadge GetEventByMM error")
			return nil, fmt.Errorf("[Service] GetBadge GetEventByMM error: %v", err)
		}

		location, err := s.artLocListRepository.GetLocationByPlaceId(event.PlaceId)
		if err != nil {
			s.logger.Debug("[Service] GetBadge GetLocationByPlaceId error")
			return nil, fmt.Errorf("[Service] GetBadge GetLocationByPlaceId error: %v", err)
		}

		location.Details = struct {
			Event domain.ArtEvent `json:"event" form:"event"`
		}{
			Event: *event,
		}

		return &location, nil
	}

	return &resp, nil
}
