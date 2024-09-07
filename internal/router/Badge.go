package router

import (
	"github.com/gin-gonic/gin"
	"go-cleanarch/internal/controller"
	"go.uber.org/zap"
)

func addBadgeGroup(router *gin.Engine, logger *zap.Logger, badgeController *controller.BadgeController) {
	BadgeGroup := router.Group("/api/v1/beacon")
	{
		BadgeGroup.GET("", badgeController.GetBadge)
	}
	logger.Info("Badge router Group added")
}
