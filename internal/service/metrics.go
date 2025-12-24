package service

import "github.com/Dmitry145528/go-metrics-tpl.git/internal/repository"

type MetricsService struct {
	repo *repository.MemStorage
}

func NewMetricsService(repo *repository.MemStorage) *MetricsService {
	return &MetricsService{
		repo: repo,
	}
}

func (s *MetricsService) UpdateGauge(name string, value float64) {
	s.repo.UpdateGauge(name, value)
}

func (s *MetricsService) UpdateCounter(name string, delta int64) {
	s.repo.UpdateCounter(name, delta)
}
