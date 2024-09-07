package service

import (
	"go-cleanarch/pkg/domain"

	"go.uber.org/zap"
)

type BadgeService struct {
	badgeRepository domain.BadgeRepository
	logger          *zap.Logger
}

func NewBadgeService(badgeRepository domain.BadgeRepository, logger *zap.Logger) *BadgeService {
	return &BadgeService{
		badgeRepository: badgeRepository,
		logger:          logger,
	}
}

func (s *BadgeService) GetBadge(mm uint32, id int) (*domain.LostItem, error) {
	s.logger.Debug("[Service] GetBadge")

	var resp *domain.Location
	resp.MajorMinor = mm

	// use MM to get locationId, sublocationId
	locationId, sublocationId, err := s.badgeRepository.GetLocationByMM(mm)
	if err != nil {
		s.logger.Debug("[Service] GetBadge GetLocationByMM error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp.LocationName, err = s.badgeRepository.GetNameByLocation(locationId)
	if err != nil {
		s.logger.Debug("[Service] GetBadge GetNameByLocation error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// get main badge info by locationId
	resp.MainBadge, err = s.badgeRepository.GetMainBadgeByLocationId(locationId)
	if err != nil {
		s.logger.Debug("[Service] GetBadge GetMainBadgeByLocationId error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	// record current sublocation to visit_log
	s.badgeRepository.AddVisitLog(mm, locationId, sublocationId)
	// get all sublocation info by locationId
	subBadgesFromDB := getSubLocListByLocId(locationId)

	// get user's visited record
	visitList := s.badgeRepository.GetVisitedSubLocIdsByUserLocInfo(id, locationId)
	
	// compare sublocation info with user's visited record
	visitedMap := make(map[int]bool)
	for _, visit := range visitList {
		visitedMap[visit] = true
	}

	var respSubBadges []domain.SubBadge
	countProgress := 0
	for _, subBadge := range subBadges {
		var badge domain.Badge{
			IconPath:    subBadge.IconPath,
			Description: subBadge.Description,
		}
		
		if visitedMap[subBadge.SubId] {
			badge.Aquired = true
			countProgress++
		} else {
			badge.Aquired = false
		}

		var subBadgeTmp domain.SubBadge{
			Badge:		badge,
			SubId:      subBadge.SubId,
		}

		respSubBadges = append(respSubBadges, subBadgeTmp)
	}

	resp.SubBadges = respSubBadges
	resp.Progress = countProgress
	resp.TotalProgress = s.badgeRepository.GetNumOfSubLocByLocId(locationId)

	return resp, nil
}
