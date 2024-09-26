package models

type Event struct {
	Id          int    `json:"id"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
	LocationId  *int   `json:"locationId" db:"location_id"`
	CreatedBy   *int   `json:"createBy" db:"created_by"`
}

type EventJoinLocation struct {
	Id          int    `json:"id"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Location    string `json:"location"`
	CreatedBy   *int   `json:"createBy" db:"created_by"`
}
