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

type LocationController struct {
	LocationService *service.LocationService
	logger          *zap.Logger
}

func NewBadgeController(locationService *service.LocationService, logger *zap.Logger) *LocationController {
	return &LocationController{
		LocationService: locationService,
		logger:          logger,
	}
}

func (bc *LocationController) GetBadge(c *gin.Context) {
	bc.logger.Info("enter controller")
	mm, err := strconv.ParseUint(c.Query("mm"), 10, 32)
	umm := domain.MajorMinor(uint32(mm))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	badge, err := bc.LocationService.GetBadge(umm, id)

	if errors.Is(err, domain.ErrNotFound) {
		c.Status(http.StatusNotFound)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, badge)
}
