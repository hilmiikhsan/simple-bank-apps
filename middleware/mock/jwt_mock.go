// Code generated by MockGen. DO NOT EDIT.
// Source: jwt.go

// Package mock_middleware is a generated GoMock package.
package mock_middleware

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	middleware "github.com/simple-bank-apps/middleware"
)

// MockJWT is a mock of JWT interface.
type MockJWT struct {
	ctrl     *gomock.Controller
	recorder *MockJWTMockRecorder
}

// MockJWTMockRecorder is the mock recorder for MockJWT.
type MockJWTMockRecorder struct {
	mock *MockJWT
}

// NewMockJWT creates a new mock instance.
func NewMockJWT(ctrl *gomock.Controller) *MockJWT {
	mock := &MockJWT{ctrl: ctrl}
	mock.recorder = &MockJWTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWT) EXPECT() *MockJWTMockRecorder {
	return m.recorder
}

// DeleteTokenFromRedis mocks base method.
func (m *MockJWT) DeleteTokenFromRedis(ctx context.Context, id, authKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTokenFromRedis", ctx, id, authKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTokenFromRedis indicates an expected call of DeleteTokenFromRedis.
func (mr *MockJWTMockRecorder) DeleteTokenFromRedis(ctx, id, authKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTokenFromRedis", reflect.TypeOf((*MockJWT)(nil).DeleteTokenFromRedis), ctx, id, authKey)
}

// ExtractJWTClaims mocks base method.
func (m *MockJWT) ExtractJWTClaims(ctx context.Context, token string) (*middleware.JWTClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractJWTClaims", ctx, token)
	ret0, _ := ret[0].(*middleware.JWTClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractJWTClaims indicates an expected call of ExtractJWTClaims.
func (mr *MockJWTMockRecorder) ExtractJWTClaims(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractJWTClaims", reflect.TypeOf((*MockJWT)(nil).ExtractJWTClaims), ctx, token)
}

// GenerateJWTToken mocks base method.
func (m *MockJWT) GenerateJWTToken(ctx context.Context, key string, req middleware.JWTRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateJWTToken", ctx, key, req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateJWTToken indicates an expected call of GenerateJWTToken.
func (mr *MockJWTMockRecorder) GenerateJWTToken(ctx, key, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateJWTToken", reflect.TypeOf((*MockJWT)(nil).GenerateJWTToken), ctx, key, req)
}

// GetTokenFromRedis mocks base method.
func (m *MockJWT) GetTokenFromRedis(ctx context.Context, id, authKey string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenFromRedis", ctx, id, authKey)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTokenFromRedis indicates an expected call of GetTokenFromRedis.
func (mr *MockJWTMockRecorder) GetTokenFromRedis(ctx, id, authKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenFromRedis", reflect.TypeOf((*MockJWT)(nil).GetTokenFromRedis), ctx, id, authKey)
}

// SaveTokenToRedis mocks base method.
func (m *MockJWT) SaveTokenToRedis(ctx context.Context, timeLimit int, token, id, authKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTokenToRedis", ctx, timeLimit, token, id, authKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTokenToRedis indicates an expected call of SaveTokenToRedis.
func (mr *MockJWTMockRecorder) SaveTokenToRedis(ctx, timeLimit, token, id, authKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTokenToRedis", reflect.TypeOf((*MockJWT)(nil).SaveTokenToRedis), ctx, timeLimit, token, id, authKey)
}

// ValidateTokenIssuer mocks base method.
func (m *MockJWT) ValidateTokenIssuer(claims *middleware.JWTClaims) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTokenIssuer", claims)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateTokenIssuer indicates an expected call of ValidateTokenIssuer.
func (mr *MockJWTMockRecorder) ValidateTokenIssuer(claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTokenIssuer", reflect.TypeOf((*MockJWT)(nil).ValidateTokenIssuer), claims)
}
