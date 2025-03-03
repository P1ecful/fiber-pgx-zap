package service

import (
	"context"
	"log/slog"
	"srv-tmpl/internal/models/dto"
	"srv-tmpl/internal/storage"
)

type SrvTmpl interface {
	MethodFirst(ctx context.Context) *dto.Model
	MethodSecond(ctx context.Context) string
}

type Service struct {
	logger  *slog.Logger
	storage storage.Storage
}

func NewService(logger *slog.Logger, storage storage.Storage) *Service {
	return &Service{
		logger:  logger,
		storage: storage,
	}
}

func (s *Service) MethodFirst(_ context.Context) *dto.Model {
	return &dto.Model{}
}
func (s *Service) MethodSecond(_ context.Context) string { return "" }
