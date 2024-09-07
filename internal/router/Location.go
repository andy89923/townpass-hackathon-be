package router

import (
	"go-cleanarch/internal/controller"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func addBadgeGroup(router *gin.Engine, logger *zap.Logger, badgeController *controller.LocationController) {
	BadgeGroup := router.Group("/api/v1/beacon")
	{
		BadgeGroup.GET("", badgeController.GetBadge)
	}
	logger.Info("Badge router Group added")
}
