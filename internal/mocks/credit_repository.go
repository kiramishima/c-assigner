// Code generated by MockGen. DO NOT EDIT.
// Source: .\interfaces\credit_repository.go
//
// Generated by this command:
//
//	mockgen -source .\interfaces\credit_repository.go -destination .\internal\mocks\credit_repository.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models2 "kiramishima/credit_assigner/internal/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCreditRepository is a mock of CreditRepository interface.
type MockCreditRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCreditRepositoryMockRecorder
}

// MockCreditRepositoryMockRecorder is the mock recorder for MockCreditRepository.
type MockCreditRepositoryMockRecorder struct {
	mock *MockCreditRepository
}

// NewMockCreditRepository creates a new mock instance.
func NewMockCreditRepository(ctrl *gomock.Controller) *MockCreditRepository {
	mock := &MockCreditRepository{ctrl: ctrl}
	mock.recorder = &MockCreditRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreditRepository) EXPECT() *MockCreditRepositoryMockRecorder {
	return m.recorder
}

// RegisterAssign mocks base method.
func (m *MockCreditRepository) RegisterAssign(ctx context.Context, data *models2.Credit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAssign", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterAssign indicates an expected call of RegisterAssign.
func (mr *MockCreditRepositoryMockRecorder) RegisterAssign(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAssign", reflect.TypeOf((*MockCreditRepository)(nil).RegisterAssign), ctx, data)
}

// SumUp mocks base method.
func (m *MockCreditRepository) SumUp(ctx context.Context) (*models2.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SumUp", ctx)
	ret0, _ := ret[0].(*models2.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SumUp indicates an expected call of SumUp.
func (mr *MockCreditRepositoryMockRecorder) SumUp(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SumUp", reflect.TypeOf((*MockCreditRepository)(nil).SumUp), ctx)
}
