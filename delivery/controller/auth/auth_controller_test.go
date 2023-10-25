package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/simple-bank-apps/dto"
	mockusecase "github.com/simple-bank-apps/usecase/auth/mock"
	"github.com/simple-bank-apps/utils"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUsecase := mockusecase.NewMockAuthUsecase(ctrl)
	mockAuthController := &AuthController{
		authUsecase: mockAuthUsecase,
	}

	password, _ := utils.HashPassword("12345678")

	t.Run("success", func(t *testing.T) {
		mockAuth := dto.Request{
			Username: "User Test",
			Password: password,
		}

		mockAuthUsecase.EXPECT().Register(gomock.Any(), mockAuth).Return(nil)

		reqBody, _ := json.Marshal(mockAuth)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
		recorder := httptest.NewRecorder()

		gin.SetMode(gin.TestMode)

		router := gin.Default()
		router.POST("/register", mockAuthController.Register)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code)
	})

	t.Run("error", func(t *testing.T) {
		mockAuth := dto.Request{
			Username: "User Test",
			Password: "12345",
		}

		reqBody, _ := json.Marshal(mockAuth)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
		recorder := httptest.NewRecorder()

		gin.SetMode(gin.TestMode)
		router := gin.Default()

		router.POST("/register", mockAuthController.Register)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUsecase := mockusecase.NewMockAuthUsecase(ctrl)
	mockAuthController := &AuthController{
		authUsecase: mockAuthUsecase,
	}

	password, _ := utils.HashPassword("12345678")

	t.Run("success", func(t *testing.T) {
		mockAuth := dto.Request{
			Username: "User Test",
			Password: password,
		}

		mockResponse := dto.LoginResponse{
			ID:       "1",
			Token:    "token",
			Username: "User Test",
			ExpireAt: time.Now().Add(time.Duration(1) * time.Hour),
		}

		mockAuthUsecase.EXPECT().Login(gomock.Any(), mockAuth).Return(mockResponse, nil)

		reqBody, _ := json.Marshal(mockAuth)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
		recorder := httptest.NewRecorder()

		gin.SetMode(gin.TestMode)

		router := gin.Default()
		router.POST("/login", mockAuthController.Login)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("error", func(t *testing.T) {
		mockAuth := dto.Request{
			Username: "User Test",
			Password: "12345",
		}

		reqBody, _ := json.Marshal(mockAuth)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
		recorder := httptest.NewRecorder()

		gin.SetMode(gin.TestMode)
		router := gin.Default()

		router.POST("/login", mockAuthController.Login)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}
