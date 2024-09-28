package services

import (
	"context"
	"fmt"
	db "payment/db/sqlc"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentTransactionService struct {
	*db.Queries
	dbConn *pgxpool.Conn
}

func NewPaymentTransactionService(dbConn *pgxpool.Conn) *PaymentTransactionService{
	return &PaymentTransactionService{
		Queries: db.New(dbConn),
		dbConn: dbConn,
	}
}

type UserBalanceDto struct {
	UsacAccountNo *string    `json:"usac_account_no"`
	UsacBalance   float32 	 `json:"usac_balance"`
}

type TransferSendReceiveInfo struct {
	Sender *db.PaymentTransaction
	Receiver *db.PaymentTransaction
}

func(transaction *PaymentTransactionService) CreateTransferTx(c context.Context,args db.TransferParams)(*TransferSendReceiveInfo, error){
	tx, err := transaction.dbConn.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(context.Background())

	qtx := transaction.Queries.WithTx(tx)

	// ------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------
	// Send Transfer 
	argsSend := &db.TransferSendParams{
		PatrxNo: args.PatrxNo,
		PatrxCreatedOn: args.PatrxCreatedOn,
		PatrxCredit: args.PatrxCredit,
		PatrxNotes: args.PatrxNotes,
		PatrxAcctnoTo: args.PatrxAcctnoTo,
		PatrxTratyID: args.PatrxTratyID,
	}

	send, err := qtx.TransferSend(c, *argsSend)
	if err != nil {
		return nil, err
	}

	// Find User Account by No
	accSender, err := qtx.FindUserAccountByAccno(c, args.PatrxAcctnoFrom)
	if err != nil {
		return nil, err
	}
	
	// Update Balance for sender
	usacBalance := *accSender.UsacBalance
	paymentCredit := *args.PatrxCredit
	userBalance := usacBalance - paymentCredit

	argsSenderBalance := &db.UpdateBalanceParams{
		UsacAccountNo: args.PatrxAcctnoFrom,
		UsacBalance:   &userBalance,
	}  
	// Update Balance for sender
	qtx.UpdateBalance(c, *argsSenderBalance)
	
	// ------------------------------------------------------------------------------
	// Receive Transfer
	// Generate PatrxNo
	generatePatrxNo, err := incrementString(send.PatrxNo)
	if err != nil {
		return nil, err
	}

	// Receive Transfer
	argsReceive := &db.TransferReceiveParams{
		PatrxNo: generatePatrxNo,
		PatrxCreatedOn: args.PatrxCreatedOn,
		PatrxDebet: args.PatrxDebet,
		PatrxNotes: args.PatrxNotes,
		PatrxAcctnoFrom: args.PatrxAcctnoFrom,
		PatrxTratyID: args.PatrxTratyID,
		PatrxPatrxRef: &args.PatrxNo,
	}
	
	receive, err := qtx.TransferReceive(c, *argsReceive)
	if err != nil {
		return nil, err
	}

	// Find UserAccount for Balance 
	accReceiver, err := qtx.FindUserAccountByAccno(c, args.PatrxAcctnoTo)
	if err != nil {
		return nil, err
	}

	// Update Balance for sender
	usacBalance = *accReceiver.UsacBalance
	paymentDebet := *args.PatrxDebet
	userBalance = usacBalance + paymentDebet

	argsReceiveBalance := &db.UpdateBalanceParams{
		UsacAccountNo: send.PatrxAcctnoTo,
		UsacBalance:   &userBalance,
	}  
	// Update Balance for sender
	qtx.UpdateBalance(c, *argsReceiveBalance)

	TransferInfo := &TransferSendReceiveInfo{
		Sender: send,
		Receiver: receive,
	}

	if err = tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return TransferInfo, nil
}

func incrementString(input string) (string, error) {
	// Split the input string into prefix and numeric parts
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid input format")
	}

	prefix := parts[0]
	numericPart := parts[1]

	// Parse the numeric part
	num, err := strconv.Atoi(numericPart)
	if err != nil {
		return "", fmt.Errorf("error parsing numeric part: %v", err)
	}

	// Increment the numeric part
	num++

	// Format the incremented numeric part with leading zeros
	newNumericPart := fmt.Sprintf("%05d", num)

	// Combine the prefix and new numeric part
	result := fmt.Sprintf("%s-%s", prefix, newNumericPart)

	return result, nil
}