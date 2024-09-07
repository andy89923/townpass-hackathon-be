package repository

import (
	"fmt"

	"gorm.io/driver/postgres" 
	"gorm.io/gorm"

	"go.uber.org/zap"
	"go-cleanarch/pkg/factory/config"
)

func ConnTotDB(logger *zap.Logger) (*gorm.DB) {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d ", cfg.Database.Host, cfg.Database.Username, cfg.Database.Password, cfg.Database.DbName, cfg.Database.Port)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Error connecting to database", zap.Error(err))
		panic(err)
	}
	logger.Info("Database connected successfully")

	// Migrate the schema LostItem
	if err := db.AutoMigrate(&LostItem{}); err != nil {
		logger.Error("Error migrating schema", zap.Error(err))
		panic(err)
	}
	logger.Info("Database schema migrated successfully")

	// Migrate the schema LocationTable
	if err := db.AutoMigrate(&LocationTable{}); err != nil {
		logger.Error("Error migrating schema", zap.Error(err))
		panic(err)
	}
	logger.Info("Database schema LocationTable migrated successfully")


	return db
}
