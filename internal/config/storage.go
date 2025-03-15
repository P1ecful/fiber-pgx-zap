package config

import (
	"fmt"
	"go.uber.org/zap"
)

type StorageConfig struct {
	Host     string
	Port     string
	Database string
	Password string
	Username string
	URI      string
}

func (s *StorageConfig) SetURI(logger *zap.Logger) {
	s.URI = fmt.Sprintf(
		s.URI,
		s.Username,
		s.Password,
		s.Host,
		s.Port,
		s.Database,
	)

	logger.Info("successfully set uri", zap.Any("uri", s.URI))
}

func (s *StorageConfig) GetURI() string {
	return s.URI
}
