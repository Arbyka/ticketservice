package repository

import (
	"ticketing/entity"
	"gorm.io/gorm"
)

type EventImageRepository interface {
	Create(image *entity.EventImage) error
}

type eventImageRepository struct {
	db *gorm.DB
}

func NewEventImageRepository(db *gorm.DB) EventImageRepository {
	return &eventImageRepository{db}
}

func (r *eventImageRepository) Create(image *entity.EventImage) error {
	return r.db.Create(image).Error
}
