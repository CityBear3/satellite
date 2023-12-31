// Code generated by MockGen. DO NOT EDIT.
// Source: archive.go
//
// Generated by this command:
//
//	mockgen -source=archive.go -package=mock_usecase -destination=./mock/archive.go
//
// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	entity "github.com/CityBear3/satellite/internal/domain/entity"
	usecase "github.com/CityBear3/satellite/internal/usecase"
	gomock "go.uber.org/mock/gomock"
)

// MockArchiveUseCase is a mock of ArchiveUseCase interface.
type MockArchiveUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockArchiveUseCaseMockRecorder
}

// MockArchiveUseCaseMockRecorder is the mock recorder for MockArchiveUseCase.
type MockArchiveUseCaseMockRecorder struct {
	mock *MockArchiveUseCase
}

// NewMockArchiveUseCase creates a new mock instance.
func NewMockArchiveUseCase(ctrl *gomock.Controller) *MockArchiveUseCase {
	mock := &MockArchiveUseCase{ctrl: ctrl}
	mock.recorder = &MockArchiveUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArchiveUseCase) EXPECT() *MockArchiveUseCaseMockRecorder {
	return m.recorder
}

// CreateArchive mocks base method.
func (m *MockArchiveUseCase) CreateArchive(ctx context.Context, request usecase.CreateArchiveInput, device entity.Device) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArchive", ctx, request, device)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateArchive indicates an expected call of CreateArchive.
func (mr *MockArchiveUseCaseMockRecorder) CreateArchive(ctx, request, device any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArchive", reflect.TypeOf((*MockArchiveUseCase)(nil).CreateArchive), ctx, request, device)
}

// GetArchive mocks base method.
func (m *MockArchiveUseCase) GetArchive(ctx context.Context, request usecase.GetArchiveInput, client entity.Client) (usecase.GetArchiveResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArchive", ctx, request, client)
	ret0, _ := ret[0].(usecase.GetArchiveResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArchive indicates an expected call of GetArchive.
func (mr *MockArchiveUseCaseMockRecorder) GetArchive(ctx, request, client any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArchive", reflect.TypeOf((*MockArchiveUseCase)(nil).GetArchive), ctx, request, client)
}
