package controller

import (
	"errors"
	"strconv"

	"go-cleanarch/internal/service"
	"go-cleanarch/pkg/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LostItemController struct {
	lostItemService *service.LostItemService
	logger 		*zap.Logger
}

func NewLostItemController(lostItemService *service.LostItemService, logger *zap.Logger) *LostItemController {
	return &LostItemController{
		lostItemService: lostItemService,
		logger: logger,
	}
}

func (lic *LostItemController) PostOne(c *gin.Context) {
	lic.logger.Debug("PostOne")

	var lostItem domain.LostItem
	err := c.ShouldBindJSON(&lostItem)
	if err != nil {
		lic.logger.Error("PostOne", zap.Error(err))
		c.Status(http.StatusBadRequest)
		return
	}

	newLostItem, err := lic.lostItemService.AddNewLostItem(&lostItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newLostItem)
}

func (lic *LostItemController) GetAll(c *gin.Context) {
	lostItems, err := lic.lostItemService.GetAllLostItems()
	
	if errors.Is(err, domain.ErrNotFound) {
		c.Status(http.StatusNotFound)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lostItems)
}

func (lic *LostItemController) GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lostItem, err := lic.lostItemService.GetLostItemById(uint(id))
	if errors.Is(err, domain.ErrNotFound) {
		c.Status(http.StatusNotFound)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lostItem)
}

func (lic *LostItemController) UpdateOne(c *gin.Context) {
	var lostItem domain.LostItem
	err := c.ShouldBindJSON(&lostItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = lic.lostItemService.UpdateLostItem(&lostItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (lic *LostItemController) DeleteOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = lic.lostItemService.DeleteLostItem(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}