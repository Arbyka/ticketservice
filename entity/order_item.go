package entity

type OrderItem struct {
	ID        	uint    `gorm:"primaryKey"`
	TicketID   	uint
	EventID		uint
	Event   	Event
	Quantity  	int
	Price     	float64
}

