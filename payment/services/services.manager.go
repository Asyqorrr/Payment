package services

import "github.com/jackc/pgx/v5/pgxpool"

type ServiceManager struct {
	*EntityService
	*UserService
	*PaymentTransactionService
}

func NewServiceManager(dbConn *pgxpool.Conn) *ServiceManager {
	return &ServiceManager{
		EntityService: NewEntityService(dbConn),
		UserService: NewUserService(dbConn),
		PaymentTransactionService: NewPaymentTransactionService(dbConn),
	}
}