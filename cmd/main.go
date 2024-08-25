package main

import (
	"fmt"

	"gorm.io/driver/postgres" 
	"gorm.io/gorm"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/spf13/viper"

	"go-cleanarch/pkg/domain"
	"go-cleanarch/internal/repository"
	"go-cleanarch/internal/service"
	"go-cleanarch/internal/router"
)

type Config struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Network string `mapstructure:"network"`
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	DbName string `mapstructure:"dbName"`
}

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

func initDB(cfg Config, logger *zap.Logger) (*gorm.DB) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d ", cfg.Host, cfg.Username, cfg.Password, cfg.DbName, cfg.Port)
	

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Error connecting to database", zap.Error(err))
		panic(err)
	}
	logger.Info("Database connected successfully")

	// Migrate the schema
	if err := db.AutoMigrate(&domain.LostItem{}); err != nil {
		logger.Error("Error migrating schema", zap.Error(err))
		panic(err)
	}
	logger.Info("Database schema migrated successfully")

	return db
}

func readConfig(config *Config, path string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")
	viper.SetDefault("DB_DSN", "") // Default value for DB_DSN
	viper.AutomaticEnv()

	if path != "" {
		viper.SetConfigFile(path)
	}

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("[Config] Config file not found. Reading from environment variables.")
		} else {
			fmt.Println("[Config] Error reading config file", "err", err)
			panic(err)
		}
	}

	viper.Unmarshal(config)
	fmt.Printf("[Config] Config read config: %+v\n", config)
}

func main() {
	var config Config
	readConfig(&config, "")
	logger := initLogger()

	//initialize database
	db := initDB(config, logger)
	lostItemRepository := repository.NewPostgresLostItemRepository(db, logger)

	lostItemService := service.NewLostItemService(lostItemRepository, logger)
	services := service.AppService{
		LostItemService: lostItemService,
	}
	//initialize router


	router.NewRouter(logger, services).Run(":8080")
	logger.Info("Starting application")
}
