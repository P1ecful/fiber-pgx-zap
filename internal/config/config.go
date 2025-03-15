package config

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

type Config struct {
	Service ServiceConfig
	Storage StorageConfig
}

type ServiceConfig struct {
	Port string
}

func LoadConfig(path string, logger *zap.Logger) *Config {
	if err := godotenv.Load(path); err != nil {
		logger.Debug("config file does not exist", zap.Any("path", path))

		return nil
	}

	var cfg = &Config{
		Service: ServiceConfig{
			Port: os.Getenv("SERVICE_PORT"),
		},
		Storage: StorageConfig{
			Host:     os.Getenv("STORAGE_HOST"),
			Port:     os.Getenv("STORAGE_PORT"),
			Username: os.Getenv("STORAGE_USERNAME"),
			Password: os.Getenv("STORAGE_PASSWORD"),
			Database: os.Getenv("STORAGE_DATABASE"),
			URI:      os.Getenv("STORAGE_URI"),
		},
	}

	logger.Info("successfully loaded config", zap.Any("config", cfg))
	return cfg
}
