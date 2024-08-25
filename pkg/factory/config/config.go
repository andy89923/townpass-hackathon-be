package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Network string `mapstructure:"network"`
		Host string `mapstructure:"host"`
		Port int `mapstructure:"port"`
		DbName string `mapstructure:"dbName"`
	}
	
	Server struct {
		Port int `mapstructure:"port"`
	}
}

var (
	config *Config
)

func GetConfig() *Config {
	return config
}

func ReadConfig(path string) {
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

	config = &Config{}
	viper.Unmarshal(config)
	fmt.Printf("[Config] Config read config: %+v\n", config)
}