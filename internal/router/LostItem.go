package router

import (

	"go-cleanarch/internal/controller"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func addLostItemGroup(router *gin.Engine, logger *zap.Logger, lostItemController *controller.LostItemController) {
	LostItemGroup := router.Group("/lost-items")
	{
		LostItemGroup.GET("", lostItemController.GetAll)
		LostItemGroup.POST("", lostItemController.PostOne)
	}
	logger.Info("LostItem router Group added")
}
