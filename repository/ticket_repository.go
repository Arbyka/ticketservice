package repository

import (
	"gorm.io/gorm"
	"ticketing/entity"
)

type TicketRepository interface {
	FindAll() ([]entity.Ticket, error)
	FindByID(id uint) (entity.Ticket, error)
	Create(ticket entity.Ticket) (entity.Ticket, error)
	UpdateStatus(id uint, status string) error
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) FindAll() ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	err := r.db.Preload("User").Preload("OrderItems").Find(&tickets).Error
	return tickets, err
}

func (r *ticketRepository) FindByID(id uint) (entity.Ticket, error) {
	var ticket entity.Ticket
	err := r.db.Preload("User").Preload("OrderItems").First(&ticket, id).Error
	return ticket, err
}

func (r *ticketRepository) Create(ticket entity.Ticket) (entity.Ticket, error) {
	err := r.db.Create(&ticket).Error
	return ticket, err
}

func (r *ticketRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&entity.Ticket{}).Where("id = ?", id).Update("status", status).Error
}
