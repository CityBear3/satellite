// Code generated by MockGen. DO NOT EDIT.
// Source: evnet.go
//
// Generated by this command:
//
//	mockgen -source=evnet.go -package=mock_repository -destination=../../adaptor/repository/mock/evnet.go
//
// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	entity "github.com/CityBear3/satellite/internal/domain/entity"
	primitive "github.com/CityBear3/satellite/internal/domain/primitive"
	repository "github.com/CityBear3/satellite/internal/domain/repository"
	gomock "go.uber.org/mock/gomock"
)

// MockIEventRepository is a mock of IEventRepository interface.
type MockIEventRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIEventRepositoryMockRecorder
}

// MockIEventRepositoryMockRecorder is the mock recorder for MockIEventRepository.
type MockIEventRepositoryMockRecorder struct {
	mock *MockIEventRepository
}

// NewMockIEventRepository creates a new mock instance.
func NewMockIEventRepository(ctrl *gomock.Controller) *MockIEventRepository {
	mock := &MockIEventRepository{ctrl: ctrl}
	mock.recorder = &MockIEventRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEventRepository) EXPECT() *MockIEventRepositoryMockRecorder {
	return m.recorder
}

// GetArchiveEvent mocks base method.
func (m *MockIEventRepository) GetArchiveEvent(ctx context.Context, archiveEventID primitive.ID) (entity.ArchiveEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArchiveEvent", ctx, archiveEventID)
	ret0, _ := ret[0].(entity.ArchiveEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArchiveEvent indicates an expected call of GetArchiveEvent.
func (mr *MockIEventRepositoryMockRecorder) GetArchiveEvent(ctx, archiveEventID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArchiveEvent", reflect.TypeOf((*MockIEventRepository)(nil).GetArchiveEvent), ctx, archiveEventID)
}

// SaveArchiveEvent mocks base method.
func (m *MockIEventRepository) SaveArchiveEvent(ctx context.Context, tx repository.ITx, archiveEvent entity.ArchiveEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveArchiveEvent", ctx, tx, archiveEvent)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveArchiveEvent indicates an expected call of SaveArchiveEvent.
func (mr *MockIEventRepositoryMockRecorder) SaveArchiveEvent(ctx, tx, archiveEvent any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveArchiveEvent", reflect.TypeOf((*MockIEventRepository)(nil).SaveArchiveEvent), ctx, tx, archiveEvent)
}
