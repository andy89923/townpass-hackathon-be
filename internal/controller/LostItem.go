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

	newLostItem, err := lic.lostItemService.AddNewLostItem(&lostItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newLostItem)
}

func (lic *LostItemController) GetAll(c *gin.Context) {
	// todos, err := lic.todoService.GetAllTodos()
	// if errors.Is(err, domain.ErrNotFound) {
	// 	c.Status(http.StatusNotFound)
	// 	return
	// } else if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, todos)

	c.Status(http.StatusNotImplemented)
}
