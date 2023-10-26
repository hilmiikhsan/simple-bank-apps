package bank

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/simple-bank-apps/dto"
	mockusecase "github.com/simple-bank-apps/usecase/bank/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetListBank(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankUsecase := mockusecase.NewMockBankUsecase(ctrl)
	mockBankController := &BankController{
		bankUsecase: mockBankUsecase,
	}

	t.Run("success", func(t *testing.T) {
		mockListBank := []dto.BankResponse{
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
		}

		mockBankUsecase.EXPECT().GetListBank(gomock.Any()).Return(mockListBank, nil)

		req, _ := http.NewRequest("GET", "/bank", nil)
		recorder := httptest.NewRecorder()

		gin.SetMode(gin.TestMode)
		router := gin.Default()

		router.GET("/bank", mockBankController.GetListBank)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response struct {
			Data []dto.BankResponse `json:"data"`
		}

		err := json.NewDecoder(recorder.Body).Decode(&response)
		if err != nil {
			t.Errorf("Error decoding response body: %v", err)
		}

		assert.Len(t, response.Data, len(mockListBank))
	})

	t.Run("error", func(t *testing.T) {
		mockBankUsecase.EXPECT().GetListBank(gomock.Any()).Return(nil, assert.AnError)

		req, _ := http.NewRequest("GET", "/bank", nil)
		recorder := httptest.NewRecorder()

		gin.SetMode(gin.TestMode)
		router := gin.Default()

		router.GET("/bank", mockBankController.GetListBank)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})
}
