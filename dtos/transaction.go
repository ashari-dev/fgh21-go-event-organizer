package dtos

type FormTransaction struct {
	EventId         int   `form:"eventId"`
	SectionId       []int `form:"sectionId[]"`
	TicketQty       []int `form:"ticketQty[]"`
	PaymentMethodId int   `form:"paymentMethodId"`
}