package bank

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simple-bank-apps/dto"
	"github.com/simple-bank-apps/model"
	mockBank "github.com/simple-bank-apps/repository/bank/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetListBank_Success(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankRepo := mockBank.NewMockBankRepository(ctrl)

	bankUC := NewBankUsecase(mockBankRepo)

	testCases := []struct {
		testName         string
		mockBankList     []model.Bank
		mockListErr      error
		expectedResponse []dto.BankResponse
		expectedErr      error
	}{
		{
			testName: "Success get list bank",
			mockBankList: []model.Bank{
				{
					ID:            "1",
					Name:          "Bank Test",
					AccountNumber: "1234567890",
					AccountName:   "Account Name",
				},
				{
					ID:            "2",
					Name:          "Bank Test 2",
					AccountNumber: "1234567891",
					AccountName:   "Account Name 2",
				},
			},
			mockListErr: nil,
			expectedResponse: []dto.BankResponse{
				{
					ID:            "1",
					Name:          "Bank Test",
					AccountName:   "Account Name",
					AccountNumber: "1234567890",
				},
				{
					ID:            "2",
					Name:          "Bank Test 2",
					AccountName:   "Account Name 2",
					AccountNumber: "1234567891",
				},
			},
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			mockBankRepo.EXPECT().List(ctx).Return(testCase.mockBankList, testCase.mockListErr)

			banks, err := bankUC.GetListBank(ctx)

			assert.Equal(t, testCase.expectedErr, err)
			assert.Equal(t, testCase.expectedResponse, banks)
		})
	}
}

func TestGetListBank_Failed(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankRepo := mockBank.NewMockBankRepository(ctrl)

	bankUC := NewBankUsecase(mockBankRepo)

	testCases := []struct {
		testName         string
		mockBankList     []model.Bank
		mockListErr      error
		expectedResponse []dto.BankResponse
		expectedErr      error
	}{
		{
			testName:         "Failed get list bank",
			mockBankList:     nil,
			mockListErr:      assert.AnError,
			expectedResponse: nil,
			expectedErr:      assert.AnError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			mockBankRepo.EXPECT().List(ctx).Return(testCase.mockBankList, testCase.mockListErr)

			banks, err := bankUC.GetListBank(ctx)

			assert.Equal(t, testCase.expectedErr, err)
			assert.Equal(t, testCase.expectedResponse, banks)
		})
	}
}
