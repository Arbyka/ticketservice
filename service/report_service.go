package service

import "ticketing/repository"

type ReportsService interface {
	GetSummary() (int64, float64, error)
	GetEventReport(eventID uint) (int64, float64, error)
}

type reportsService struct {
	repo repository.ReportsRepository
}

func NewReportsService(repo repository.ReportsRepository) ReportsService {
	return &reportsService{repo}
}

func (s *reportsService) GetSummary() (int64, float64, error) {
	return s.repo.GetSummaryReport()
}

func (s *reportsService) GetEventReport(eventID uint) (int64, float64, error) {
	return s.repo.GetEventReport(eventID)
}
