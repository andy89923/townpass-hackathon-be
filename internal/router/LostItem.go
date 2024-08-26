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
		LostItemGroup.GET("/:id", lostItemController.GetOne)
		LostItemGroup.PATCH("/:id", lostItemController.UpdateOne)
		LostItemGroup.DELETE("/:id", lostItemController.DeleteOne)
	}
	logger.Info("LostItem router Group added")
}
