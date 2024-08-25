package router

import (
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
	"go-cleanarch/internal/controller"
	"go-cleanarch/internal/service"
)

func NewRouter(logger *zap.Logger,  services service.AppService) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	lostItemController := controller.NewLostItemController(services.LostItemService)

	addLostItemGroup(router, logger, lostItemController)

	return router
}
