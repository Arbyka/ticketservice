package repository

import (
	"ticketing/entity"
	"gorm.io/gorm"
)

type ReportsRepository interface {
	GetSummaryReport() (totalTickets int64, totalRevenue float64, err error)
	GetEventReport(eventID uint) (ticketsSold int64, totalRevenue float64, err error)
}

type reportsRepository struct {
	db *gorm.DB
}

func NewReportsRepository(db *gorm.DB) ReportsRepository {
	return &reportsRepository{db}
}

func (r *reportsRepository) GetSummaryReport() (int64, float64, error) {
	var totalTickets int64
	var totalRevenue float64

	err := r.db.Model(&entity.OrderItem{}).
		Count(&totalTickets).
		Select("SUM(price * quantity)").Scan(&totalRevenue).Error

	return totalTickets, totalRevenue, err
}

func (r *reportsRepository) GetEventReport(eventID uint) (int64, float64, error) {
	var ticketsSold int64
	var totalRevenue float64

	err := r.db.Model(&entity.OrderItem{}).
		Where("event_id = ?", eventID).
		Count(&ticketsSold).
		Select("SUM(price * quantity)").Scan(&totalRevenue).Error

	return ticketsSold, totalRevenue, err
}
