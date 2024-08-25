package main

import (
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go-cleanarch/pkg/factory/config"
	"go-cleanarch/internal/repository"
	"go-cleanarch/internal/service"
	"go-cleanarch/internal/router"
)



func initLogger() *zap.Logger {
	logger, _ := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:      []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
				MessageKey:   "message",
				LevelKey:     "level",
				TimeKey:      "time",
				NameKey:      "logger",
				CallerKey:    "caller",
				EncodeTime:    zapcore.ISO8601TimeEncoder,
				EncodeLevel:   zapcore.LowercaseColorLevelEncoder, // can define by any self-defined function
				EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()
	
	defer logger.Sync()

	logger.Info("Logger initialized")
	return logger
}




func main() {
	config.ReadConfig( "")
	logger := initLogger()

	//initialize database
	db := repository.ConnTotDB(logger)
	lostItemRepository := repository.NewPostgresLostItemRepository(db, logger)

	lostItemService := service.NewLostItemService(lostItemRepository, logger)
	services := service.AppService{
		LostItemService: lostItemService,
	}
	//initialize router

	port := ":" + strconv.Itoa(config.GetConfig().Server.Port)
	router.NewRouter(logger, services).Run(port)
	logger.Info("Starting application")
}
