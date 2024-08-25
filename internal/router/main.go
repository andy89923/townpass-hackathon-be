package router

import (
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func NewRouter(logger *zap.Logger ) *gin.Engine {
	router := gin.Default()

	// addTodoGroup(router)
	addLostItemGroup(router, logger)

	return router
}
