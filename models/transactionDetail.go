package models

type TransactionDetails struct {
	Id            int `json:"id"`
	TransactionId int `json:"transactionId" db:"transaction_id"`
	SectionId     int `json:"sectionId" db:"section_id"`
	TicketQty     int `json:"ticketQTy" db:"ticket_qty"`
}
