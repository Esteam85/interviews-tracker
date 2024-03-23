// Code generated by MockGen. DO NOT EDIT.
// Source: process/domain/process.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	"context"
	reflect "reflect"

	"github.com/esteam85/interviews-tracker/process/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockRepository) Save(ctx context.Context, process *domain.Process) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", process)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockRepositoryMockRecorder) Save(ctx context.Context, process interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepository)(nil).Save), process)
}
