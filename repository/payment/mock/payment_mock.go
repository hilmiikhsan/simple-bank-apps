// Code generated by MockGen. DO NOT EDIT.
// Source: payment_repository_interface.go

// Package mock_payment is a generated GoMock package.
package mock_payment

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/simple-bank-apps/model"
)

// MockPaymentRepository is a mock of PaymentRepository interface.
type MockPaymentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentRepositoryMockRecorder
}

// MockPaymentRepositoryMockRecorder is the mock recorder for MockPaymentRepository.
type MockPaymentRepositoryMockRecorder struct {
	mock *MockPaymentRepository
}

// NewMockPaymentRepository creates a new mock instance.
func NewMockPaymentRepository(ctrl *gomock.Controller) *MockPaymentRepository {
	mock := &MockPaymentRepository{ctrl: ctrl}
	mock.recorder = &MockPaymentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentRepository) EXPECT() *MockPaymentRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPaymentRepository) Create(ctx context.Context, paymentModel model.Payment, customerModel model.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, paymentModel, customerModel)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPaymentRepositoryMockRecorder) Create(ctx, paymentModel, customerModel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPaymentRepository)(nil).Create), ctx, paymentModel, customerModel)
}
