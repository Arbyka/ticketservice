package entity

type Report struct {
	ID           	uint    `gorm:"primaryKey"`
	EventID    		uint
	Event      		Event
	TotalSold    	int
	TotalRevenue 	float64
	Period       	string
}
