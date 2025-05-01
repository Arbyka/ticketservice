package repository

import (
	"ticketing/entity"
	"gorm.io/gorm"
)

type EventRepository interface {
	FindAll() ([]entity.Event, error)
	FindByID(id uint) (*entity.Event, error)
	Create(event *entity.Event) error
	Update(event *entity.Event) error
	Delete(id uint) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) FindAll() ([]entity.Event, error) {
	var events []entity.Event
	err := r.db.Find(&events).Error
	return events, err
}

func (r *eventRepository) FindByID(id uint) (*entity.Event, error) {
	var event entity.Event
	err := r.db.First(&event, id).Error
	return &event, err
}

func (r *eventRepository) Create(event *entity.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) Update(event *entity.Event) error {
	return r.db.Model(&entity.Event{}).Where("id = ?", event.ID).Updates(map[string]interface{}{
		"name":           event.Name,
		"description":    event.Description,
		"category":       event.Category,
		"price":          event.Price,
		"date":           event.Date,
		"location":       event.Location,
		"capacity":       event.Capacity,
		"available_seat": event.Available_seat,
		"status":         event.Status,
		"updated_at":     event.UpdatedAt,
	}).Error
}

func (r *eventRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Event{}, id).Error
}
