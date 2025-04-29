package entity

import "time"

type Event struct {
	ID          	uint      	`gorm:"primaryKey"`
	Name        	string
	Description 	string
	Category    	string
	Price       	float64
	Date			time.Time
	Location		string
	Capacity		int
	Available_seat	int
	Status			string		`gorm:"type:enum('aktif','berlangsung','selesai')"`
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
}
