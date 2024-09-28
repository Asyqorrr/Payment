package controller

import "payment/services"

type ControllerManager struct {
	*entityController
	*PaymentTransactionController
}

// constructor
func NewControllerManager(serviceManager *services.ServiceManager) *ControllerManager{
	return &ControllerManager{
		entityController: NewEntityController(*serviceManager),
		PaymentTransactionController: NewPaymentTransactionController(*serviceManager),
	}
}