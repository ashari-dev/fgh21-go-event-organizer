package models

type Transaction struct {
	Id              int `json:"id"`
	EventId         int `json:"eventId" db:"event_id"`
	PaymentMethodId int `json:"paymentMethodId" db:"payment_method_id"`
	UserId          int `json:"userId" db:"user_id"`
}

type TransactionJoin struct {
	Id            int      `json:"id"`
	FullName      string   `json:"fullName" db:"full_name"`
	EventTitle    string   `json:"eventTitle" db:"event_title"`
	Location      *int     `json:"location" db:"location_id"`
	Date          string   `json:"date"`
	PaymentMethod string   `json:"payment_method" db:"payment_method"`
	Section       []string `json:"section" db:"section_name"`
	TicketQty     []int    `json:"ticketQty" db:"ticket_qty"`
}