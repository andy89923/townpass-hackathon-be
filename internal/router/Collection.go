package router

import (
	"go-cleanarch/internal/controller"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func addCollectionGroup(router *gin.Engine, logger *zap.Logger, collectionController *controller.CollectionController) {
	CollectionGroup := router.Group("/api/v1/collections")
	{
		CollectionGroup.GET("/:id", collectionController.GetAllCollections)
	}
	logger.Info("Collection router Group added")
}
