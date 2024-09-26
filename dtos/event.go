package dtos

type FormEvent struct {
	Image       string `form:"image"`
	Title       string `form:"title"`
	Date        string `form:"date"`
	Description string `form:"description"`
	LocationId  *int   `form:"locationId" `
}

type EventCategory struct {
	EventId    int `form:"eventId"`
	CategoryId int `form:"categoryId"`
}
type SearchEvents struct {
	Search string `form:"search"`
}
