package main

import (
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go-cleanarch/internal/repository"
	"go-cleanarch/internal/router"
	"go-cleanarch/internal/service"
	"go-cleanarch/pkg/factory/config"
)

func initLogger() *zap.Logger {
	logger, _ := zap.Config{
		Encoding:    "console",
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			NameKey:      "logger",
			CallerKey:    "caller",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder, // can define by any self-defined function
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()

	defer logger.Sync()

	logger.Info("Logger initialized")
	return logger
}

func main() {
	config.ReadConfig("../config/config.yaml")
	logger := initLogger()

	//initialize database
	db := repository.ConnTotDB(logger)

	locationRepository := repository.NewPostgresLocationRepository(db, logger)
	locListRepository := repository.NewPostgresLocListRepository(db, logger)
	subLocListRepository := repository.NewPostgresSubLocListRepository(db, logger)
	visitLogRepository := repository.NewPostgresVisitLogRepository(db, logger)
	tbMapRepository := repository.NewPostgresTbMapRepository(db, logger)
	artLocListRepository := repository.NewPostgresArtLocListRepository(db, logger)
	artEventRepository := repository.NewPostgresArtEventListRepository(db, logger)

	locationService := service.NewBadgeService(locationRepository,
		locListRepository, subLocListRepository, visitLogRepository, tbMapRepository, artLocListRepository, artEventRepository, logger)

	services := service.AppService{
		LocationService: locationService,
	}

	port := ":" + strconv.Itoa(config.GetConfig().Server.Port)
	router.NewRouter(logger, services).Run(port)
	logger.Info("Starting application")
}
