package router

import (
	"go-cleanarch/internal/controller"
	// "go-cleanarch/internal/repository"
	// "go-cleanarch/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func addLostItemGroup(router *gin.Engine, logger *zap.Logger) {
	// lostItemRepo, err := repository.NewPostgresLostItemRepository()
	// if err != nil {
	// 	logger.Error("[Router] NewPostgresLostItemRepository", "err", err)
	// 	panic(err)
	// }

	// lostItemService := service.NewLostItemService(lostItemRepo)
	lostItemController := controller.NewLostItemController(nil)

	LostItemGroup := router.Group("/lost-items")
	{
		//LostItemGroup.GET("", lostItemController.GetAll)
		LostItemGroup.POST("", lostItemController.PostOne)
	}
	logger.Info("LostItem router Group added")
}
