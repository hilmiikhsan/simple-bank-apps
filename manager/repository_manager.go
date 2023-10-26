package manager

import (
	"github.com/simple-bank-apps/repository/bank"
	"github.com/simple-bank-apps/repository/customer"
	"github.com/simple-bank-apps/repository/payment"
)

type RepositoryManager interface {
	CustomerRepository() customer.CustomerRepository
	PaymentRepository() payment.PaymentRepository
	BankRepository() bank.BankRepository
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
	return payment.NewPaymentRepository(r.infra.Connect(), r.CustomerRepository())
}

func (r *repositoryManager) BankRepository() bank.BankRepository {
	return bank.NewBankRepository(r.infra.Connect())
}
