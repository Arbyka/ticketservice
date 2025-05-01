package service

import (
	"ticketing/entity"
	"ticketing/repository"
)

type EventService interface {
	GetAll() ([]entity.Event, error)
	GetByID(id uint) (*entity.Event, error)
	Create(event *entity.Event) error
	Update(id uint, updated *entity.Event) error
	Delete(id uint) error
}

type eventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) EventService {
	return &eventService{repo}
}

func (s *eventService) GetAll() ([]entity.Event, error) {
	return s.repo.FindAll()
}

func (s *eventService) GetByID(id uint) (*entity.Event, error) {
	return s.repo.FindByID(id)
}

func (s *eventService) Create(event *entity.Event) error {
	return s.repo.Create(event)
}

func (s *eventService) Update(id uint, updated *entity.Event) error {
	event, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	*event = *updated
	event.ID = id // pastikan ID tetap
	return s.repo.Update(event)
}

func (s *eventService) Delete(id uint) error {
	return s.repo.Delete(id)
}
