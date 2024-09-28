package controller

import (
	"net/http"
	db "payment/db/sqlc"
	"payment/models"
	"payment/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type PaymentTransactionController struct {
	serviceManager *services.ServiceManager
}

type SetorTransactionParamsDto struct {
	PatrxNo         string         `json:"patrx_no"`
	PatrxCreatedOn  pgtype.Date    `json:"patrx_created_on"`
	PatrxDebet      *float64 	   `json:"patrx_debet"`
	PatrxAcctnoFrom *string        `json:"patrx_acctno_from"`
	PatrxAcctnoTo   *string        `json:"patrx_acctno_to"`
}

type TransferSendParamsDto struct {
	PatrxNo         string         `json:"patrx_no"`
	PatrxCreatedOn  pgtype.Date    `json:"patrx_created_on"`
	PatrxCredit     *float64	   `json:"patrx_debet"`
	PatrxNotes      *string        `json:"patrx_notes"`
	PatrxAcctnoFrom *string        `json:"patrx_acctno_from"`
	PatrxAcctnoTo   *string        `json:"patrx_acctno_to"`
	PatrxTratyID    *int32         `json:"patrx_traty_id"`
}

type TransferReceiveParamsDto struct {
	PatrxNo         string         `json:"patrx_no"`
	PatrxCreatedOn  pgtype.Date    `json:"patrx_created_on"`
	PatrxDebet     	*float64 		`json:"patrx_credit"`
	PatrxNotes      *string        `json:"patrx_notes"`
	PatrxAcctnoFrom *string        `json:"patrx_acctno_from"`
	PatrxAcctnoTo   *string        `json:"patrx_acctno_to"`
	PatrxTratyID    *int32         `json:"patrx_traty_id"`
	PatrxPatrxRef   *string        `json:"patrx_patrx_ref"`
}

type TransferParamsDto struct {
	PatrxNo         string      `json:"patrx_no"`
	PatrxCreatedOn  pgtype.Date `json:"patrx_created_on"`
	PatrxDebet      *float64    `json:"patrx_debet"`
	PatrxCredit     *float64    `json:"patrx_credit"`
	PatrxNotes      *string     `json:"patrx_notes"`
	PatrxAcctnoFrom *string     `json:"patrx_acctno_from"`
	PatrxAcctnoTo   *string     `json:"patrx_acctno_to"`
	PatrxPatrxRef   *string     `json:"patrx_patrx_ref"`
	PatrxTratyID    *int32      `json:"patrx_traty_id"`
}

// constructor
func NewPaymentTransactionController(servicesManager services.ServiceManager) *PaymentTransactionController {
	return &PaymentTransactionController{
		serviceManager: &servicesManager,
	}
}

func (handler *PaymentTransactionController) CreatePaymentTransaction(c *gin.Context){
	var payloadSend *TransferParamsDto
	if err := c.ShouldBindJSON(&payloadSend); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.TransferParams{
		PatrxNo: payloadSend.PatrxNo,
		PatrxCreatedOn: payloadSend.PatrxCreatedOn,
		PatrxDebet: payloadSend.PatrxDebet,
		PatrxCredit: payloadSend.PatrxCredit,
		PatrxNotes: payloadSend.PatrxNotes,
		PatrxAcctnoFrom: payloadSend.PatrxAcctnoFrom, 
		PatrxAcctnoTo: payloadSend.PatrxAcctnoTo,
		PatrxPatrxRef: payloadSend.PatrxPatrxRef,
		PatrxTratyID: payloadSend.PatrxTratyID,
	}

	transferInfo, err := handler.serviceManager.PaymentTransactionService.CreateTransferTx(c, *args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusCreated, transferInfo)
}