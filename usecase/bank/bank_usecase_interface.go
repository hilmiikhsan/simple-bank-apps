package bank

import (
	"context"

	"github.com/simple-bank-apps/dto"
	"github.com/simple-bank-apps/repository/bank"
)

type BankUsecase interface {
	GetListBank(ctx context.Context) ([]dto.BankResponse, error)
}

func NewBankUsecase(bankRepo bank.BankRepository) BankUsecase {
	return &bankUsecase{
		bankRepo: bankRepo,
	}
}
