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

type BadgeController struct {
	BadgeService *service.BadgeService
	logger       *zap.Logger
}

func NewBadgeController(badgeService *service.BadgeService, logger *zap.Logger) *BadgeController {
	return &BadgeController{
		BadgeService: badgeService,
		logger:       logger,
	}
}

func (bc *BadgeController) GetBadge(c *gin.Context) {
	bc.logger.Info("enter controller")
	mm, err := strconv.ParseUint(c.Param("mm"), 10, 32)
	umm := uint32(mm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	badge, err := bc.BadgeService.GetBadge(&umm, &id)

	if errors.Is(err, domain.ErrNotFound) {
		c.Status(http.StatusNotFound)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, badge)
}
