// Code generated by MockGen. DO NOT EDIT.
// Source: archive.go
//
// Generated by this command:
//
//	mockgen -source=archive.go -package=mock_repository -destination=../../../adaptor/gateway/repository/mock/archive.go
//
// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	entity "github.com/CityBear3/satellite/internal/domain/entity"
	primitive "github.com/CityBear3/satellite/internal/domain/primitive"
	gomock "go.uber.org/mock/gomock"
)

// MockIArchiveRepository is a mock of IArchiveRepository interface.
type MockIArchiveRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIArchiveRepositoryMockRecorder
}

// MockIArchiveRepositoryMockRecorder is the mock recorder for MockIArchiveRepository.
type MockIArchiveRepositoryMockRecorder struct {
	mock *MockIArchiveRepository
}

// NewMockIArchiveRepository creates a new mock instance.
func NewMockIArchiveRepository(ctrl *gomock.Controller) *MockIArchiveRepository {
	mock := &MockIArchiveRepository{ctrl: ctrl}
	mock.recorder = &MockIArchiveRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIArchiveRepository) EXPECT() *MockIArchiveRepositoryMockRecorder {
	return m.recorder
}

// GetArchive mocks base method.
func (m *MockIArchiveRepository) GetArchive(ctx context.Context, archiveId primitive.ID) (entity.Archive, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArchive", ctx, archiveId)
	ret0, _ := ret[0].(entity.Archive)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArchive indicates an expected call of GetArchive.
func (mr *MockIArchiveRepositoryMockRecorder) GetArchive(ctx, archiveId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArchive", reflect.TypeOf((*MockIArchiveRepository)(nil).GetArchive), ctx, archiveId)
}

// GetArchiveByArchiveEventID mocks base method.
func (m *MockIArchiveRepository) GetArchiveByArchiveEventID(ctx context.Context, archiveEventID primitive.ID) (entity.Archive, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArchiveByArchiveEventID", ctx, archiveEventID)
	ret0, _ := ret[0].(entity.Archive)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArchiveByArchiveEventID indicates an expected call of GetArchiveByArchiveEventID.
func (mr *MockIArchiveRepositoryMockRecorder) GetArchiveByArchiveEventID(ctx, archiveEventID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArchiveByArchiveEventID", reflect.TypeOf((*MockIArchiveRepository)(nil).GetArchiveByArchiveEventID), ctx, archiveEventID)
}

// Save mocks base method.
func (m *MockIArchiveRepository) Save(ctx context.Context, archive entity.Archive) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, archive)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIArchiveRepositoryMockRecorder) Save(ctx, archive any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIArchiveRepository)(nil).Save), ctx, archive)
}
