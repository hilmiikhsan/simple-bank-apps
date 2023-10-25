package manager

import (
	"github.com/simple-bank-apps/repository/customer"
	"github.com/simple-bank-apps/repository/payment"
)

type RepositoryManager interface {
	CustomerRepository() customer.CustomerRepository
	PaymentRepository() payment.PaymentRepository
}

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}

type repositoryManager struct {
	infra InfraManager
}

func (r *repositoryManager) CustomerRepository() customer.CustomerRepository {
	return customer.NewCustomerRepository(r.infra.Connect())
}

func (r *repositoryManager) PaymentRepository() payment.PaymentRepository {
	return payment.NewPaymentRepository(r.infra.Connect())
}
