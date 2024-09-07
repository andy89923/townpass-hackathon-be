package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go-cleanarch/internal/controller"
	"go-cleanarch/internal/service"

	"go.uber.org/zap"
)

func NewRouter(logger *zap.Logger, services service.AppService) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Use(cors.Default())

	// lostItemController := controller.NewLostItemController(services.LostItemService, logger)
	// addLostItemGroup(router, logger, lostItemController)
	locationController := controller.NewBadgeController(services.LocationService, logger)
	addBadgeGroup(router, logger, locationController)
	return router
}
