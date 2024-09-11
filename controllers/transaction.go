package controllers

import (
	"event-organizer/dtos"
	"event-organizer/lib"
	"event-organizer/models"
	"event-organizer/repository"

	"github.com/gin-gonic/gin"
)

func CreatTransaction(c *gin.Context) {
	var formTransaction dtos.FormTransaction
	c.Bind(&formTransaction)

	trx, err := repository.CreateTransaction(models.Transaction{
		EventId:         formTransaction.EventId,
		PaymentMethodId: formTransaction.PaymentMethodId,
		UserId:          c.GetInt("userId"),
	})
	if err != nil {
		lib.HandlerBadReq(c, "transaction failed")
		return
	}

	for i := range formTransaction.SectionId {
		repository.CreateTransactionDetail(models.TransactionDetails{
			TransactionId: trx.Id,
			SectionId:     formTransaction.SectionId[i],
			TicketQty:     formTransaction.TicketQty[i],
		})
	}

	detailTrx, err := repository.GetTransaction(trx.Id)
	if err != nil {
		lib.HandlerBadReq(c, "transaction failed")
		return
	}

	lib.HandlerOK(c, "Transaction success", detailTrx, nil)
}

func GetAllTransaction(c *gin.Context) {
	userId := c.GetInt("userId")

	transactions, err := repository.GetAllTransaction(userId)
	if err != nil {
		lib.HandlerBadReq(c, "Failed to get transactions")
		return
	}
	lib.HandlerOK(c, "Transactions found", transactions, nil)
}
