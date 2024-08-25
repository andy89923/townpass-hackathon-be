package controller

import (
	// "errors"
	"go-cleanarch/internal/service"
	"go-cleanarch/pkg/domain"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
)

type LostItemController struct {
	lostItemService *service.LostItemService
}

func NewLostItemController(lostItemService *service.LostItemService) *LostItemController {
	return &LostItemController{
		lostItemService: lostItemService,
	}
}

func (lic *LostItemController) PostOne(c *gin.Context) {
	var lostItem domain.LostItem
	err := c.ShouldBindJSON(&lostItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// -------------- call service --------------
	// newLostItem, err := lic.lostItemService.AddNewLostItem(&lostItem)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusCreated, newLostItem)

	c.Status(http.StatusNotImplemented)
}

