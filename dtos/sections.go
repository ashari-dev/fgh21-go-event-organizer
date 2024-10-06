package dtos

type Sections struct {
	Name     string `form:"name" binding:"required"`
	Price    int    `form:"price" binding:"required"`
	Quantity int    `form:"quantity" binding:"required"`
	EventId  int    `form:"eventId" binding:"required"`
}
