package entity

type EventImage struct {
	ID      uint   	`gorm:"primaryKey" json:"id"`
	EventID uint   	`json:"event_id"`
	URL     string 	`json:"url"`

	Event 	Event 	`gorm:"foreignKey:EventID" json:"-"`
}
