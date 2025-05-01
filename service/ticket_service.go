package service

import (
	"errors"
	"ticketing/entity"
	"ticketing/repository"

	"gorm.io/gorm"
)

type TicketService interface {
	GetAll() ([]entity.Ticket, error)
	GetByID(id uint) (entity.Ticket, error)
	Create(ticket entity.Ticket) (entity.Ticket, error)
	Cancel(id uint) error
}

type ticketService struct {
	ticketRepo repository.TicketRepository
	eventRepo  repository.EventRepository
	db         *gorm.DB
}

func NewTicketService(ticketRepo repository.TicketRepository, eventRepo repository.EventRepository, db *gorm.DB) TicketService {
	return &ticketService{ticketRepo, eventRepo, db}
}

func (s *ticketService) GetAll() ([]entity.Ticket, error) {
	return s.ticketRepo.FindAll()
}

func (s *ticketService) GetByID(id uint) (entity.Ticket, error) {
	return s.ticketRepo.FindByID(id)
}

func (s *ticketService) Create(ticket entity.Ticket) (entity.Ticket, error) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var totalPrice float64 = 0

		for i, item := range ticket.OrderItems {
			event, err := s.eventRepo.FindByID(item.EventID)
			if err != nil {
				return err
			}
			if event.Available_seat < item.Quantity {
				return errors.New("not enough seats available for event: " + event.Name)
			}

			event.Available_seat -= item.Quantity
			if err := tx.Save(&event).Error; err != nil {
				return err
			}

			ticket.OrderItems[i].Price = float64(item.Quantity) * event.Price
			totalPrice += ticket.OrderItems[i].Price
		}

		ticket.TotalPrice = totalPrice
		ticket.Status = "tersedia"

		if err := tx.Create(&ticket).Error; err != nil {
			return err
		}
		return nil
	})

	return ticket, err
}

func (s *ticketService) Cancel(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		ticket, err := s.ticketRepo.FindByID(id)
		if err != nil {
			return err
		}
		if ticket.Status == "cancelled" {
			return errors.New("ticket already cancelled")
		}

		for _, item := range ticket.OrderItems {
			event, err := s.eventRepo.FindByID(item.EventID)
			if err != nil {
				return err
			}
			event.Available_seat += item.Quantity
			if err := tx.Save(&event).Error; err != nil {
				return err
			}
		}

		ticket.Status = "cancelled"
		return tx.Save(&ticket).Error
	})
}
