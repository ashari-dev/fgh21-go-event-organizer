package repository

import (
	"context"
	"event-organizer/lib"
	"event-organizer/models"

	"github.com/jackc/pgx/v5"
)

func CreateTransactionDetail(data models.TransactionDetails) (models.TransactionDetails, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO transaction_details (section_id, ticket_qty, transaction_id)
			VALUES ($1, $2, $3) RETURNING *`
	row, err := db.Query(context.Background(), sql, data.SectionId, data.TicketQty, data.TransactionId)
	if err != nil {
		return models.TransactionDetails{}, err
	}

	transactionDetail, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.TransactionDetails])
	if err != nil {
		return models.TransactionDetails{}, err
	}

	return transactionDetail, nil
}
