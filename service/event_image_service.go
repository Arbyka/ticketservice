package service

import (
	"ticketing/entity"
	"ticketing/repository"
)

type EventImageService interface {
	AddImage(image *entity.EventImage) error
}

type eventImageService struct {
	repo repository.EventImageRepository
}

func NewEventImageService(repo repository.EventImageRepository) EventImageService {
	return &eventImageService{repo}
}

func (s *eventImageService) AddImage(image *entity.EventImage) error {
	return s.repo.Create(image)
}
