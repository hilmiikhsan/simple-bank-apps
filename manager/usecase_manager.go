package manager

import (
	"github.com/simple-bank-apps/config"
	"github.com/simple-bank-apps/middleware"
	"github.com/simple-bank-apps/usecase/auth"
	"github.com/simple-bank-apps/usecase/bank"
	"github.com/simple-bank-apps/usecase/payment"
)

type UsecaseManager interface {
	AuthUsecase() auth.AuthUsecase
	PaymentUsecase() payment.PaymentUsecase
	BankUsecase() bank.BankUsecase
}

func NewUsecaseManager(repo RepositoryManager, jwt middleware.JWT, cfg *config.Config, logger middleware.LogMiddleware) UsecaseManager {
	return &usecaseManager{
		repo:   repo,
		jwt:    jwt,
		cfg:    cfg,
		logger: logger,
	}
}

type usecaseManager struct {
	repo   RepositoryManager
	jwt    middleware.JWT
	cfg    *config.Config
	logger middleware.LogMiddleware
}

func (u *usecaseManager) AuthUsecase() auth.AuthUsecase {
	return auth.NewAuthUsecase(u.repo.CustomerRepository(), u.jwt, u.cfg)
}

func (u *usecaseManager) PaymentUsecase() payment.PaymentUsecase {
	return payment.NewPaymentUsecase(u.repo.PaymentRepository(), u.repo.CustomerRepository(), u.repo.BankRepository())
}

func (u *usecaseManager) BankUsecase() bank.BankUsecase {
	return bank.NewBankUsecase(u.repo.BankRepository())
}
