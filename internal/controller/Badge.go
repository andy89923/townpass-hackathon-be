package controller

import (
	"go-cleanarch/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BadgeController struct {
	BadgeService *service.BadgeService
	logger       *zap.Logger
}

func NewBadgeController(badgeService *service.BadgeService, logger *zap.Logger) *BadgeController {
	return &BadgeController{
		BadgeService: badgeService,
		logger:       logger,
	}
}

func (bc *BadgeController) GetBadge(c *gin.Context) {
	bc.logger.Info("enter controller")
	// badge, err := bc.BadgeService.GetBadge()

	// if errors.Is(err, domain.ErrNotFound) {
	// 	c.Status(http.StatusNotFound)
	// 	return
	// } else if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, badge)
}
