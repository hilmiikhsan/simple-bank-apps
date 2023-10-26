package bank

import (
	"context"

	"github.com/simple-bank-apps/dto"
	"github.com/simple-bank-apps/repository/bank"
)

type bankUsecase struct {
	bankRepo bank.BankRepository
}

func (b *bankUsecase) GetListBank(ctx context.Context) ([]dto.BankResponse, error) {
	listBank, err := b.bankRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	var listBankResponse []dto.BankResponse
	for _, bank := range listBank {
		listBankResponse = append(listBankResponse, dto.BankResponse{
			ID:            bank.ID,
			Name:          bank.Name,
			AccountName:   bank.AccountName,
			AccountNumber: bank.AccountNumber,
		})
	}

	return listBankResponse, nil
}
