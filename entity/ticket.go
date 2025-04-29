package entity

import "time"

type Ticket struct {
	ID              uint         	`gorm:"primaryKey"`
	UserID          uint
	User            User
	TotalPrice      float64
	Status          string       	`gorm:"type:enum('tersedia','habis','cancelled')"`
	OrderItems      []OrderItem		`gorm:"foreignKey:TicketID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
