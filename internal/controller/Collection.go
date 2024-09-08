package controller

import (
	"errors"
	"go-cleanarch/internal/service"
	"go-cleanarch/pkg/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CollectionController struct {
	CollectionService *service.CollectionService
	logger            *zap.Logger
}

func NewCollectionController(collectionService *service.CollectionService, logger *zap.Logger) *CollectionController {
	return &CollectionController{
		CollectionService: collectionService,
		logger:            logger,
	}
}

func (cc *CollectionController) GetAllCollections(c *gin.Context) {
	cc.logger.Info("enter controller")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collections, err := cc.CollectionService.GetCollections(id)

	if errors.Is(err, domain.ErrNotFound) {
		c.Status(http.StatusNotFound)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, collections)
}
