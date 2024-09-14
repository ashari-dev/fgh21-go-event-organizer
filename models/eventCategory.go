package models

type EventCategory struct {
	Id         int `json:"id"`
	EventId    int `json:"eventId" db:"event_id"`
	CategoryId int `json:"categoryId" db:"category_id"`
}
