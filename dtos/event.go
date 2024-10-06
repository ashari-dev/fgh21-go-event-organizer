package dtos

type FormEvent struct {
	Image       string `form:"image"`
	Title       string `form:"title" binding:"required"`
	Date        string `form:"date" binding:"required"`
	Description string `form:"description" binding:"required"`
	LocationId  *int   `form:"locationId" binding:"required"`
}

type EventCategory struct {
	EventId    int `form:"eventId" binding:"required"`
	CategoryId int `form:"categoryId" binding:"required"`
}

type SearchEvents struct {
	Search string `form:"search"`
}
