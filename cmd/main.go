package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/spf13/viper"
)

type Config struct {
	Username string `mapstructure:"Username"`
	Password string `mapstructure:"Password"`
	Network string `mapstructure:"Network"`
	Server string `mapstructure:"Server"`
	Port int `mapstructure:"Port"`
	Name string `mapstructure:"Name"`
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
	readConfig(&Config{}, "")
	logger := initLogger()

	//initialize database

	//initialize service

	//initialize repository
	
	//initialize contoller

	//initialize router

	logger.Info("Starting application")
	// router.NewRouter().Run(":8080")
}
