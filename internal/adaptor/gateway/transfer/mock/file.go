// Code generated by MockGen. DO NOT EDIT.
// Source: file.go
//
// Generated by this command:
//
//	mockgen -source=file.go -package=mock_transfer -destination=../../../adaptor/gateway/transfer/mock/file.go
//
// Package mock_transfer is a generated GoMock package.
package mock_transfer

import (
	context "context"
	reflect "reflect"

	primitive "github.com/CityBear3/satellite/internal/domain/primitive"
	archive "github.com/CityBear3/satellite/internal/domain/primitive/archive"
	gomock "go.uber.org/mock/gomock"
)

// MockIFileTransfer is a mock of IFileTransfer interface.
type MockIFileTransfer struct {
	ctrl     *gomock.Controller
	recorder *MockIFileTransferMockRecorder
}

// MockIFileTransferMockRecorder is the mock recorder for MockIFileTransfer.
type MockIFileTransferMockRecorder struct {
	mock *MockIFileTransfer
}

// NewMockIFileTransfer creates a new mock instance.
func NewMockIFileTransfer(ctrl *gomock.Controller) *MockIFileTransfer {
	mock := &MockIFileTransfer{ctrl: ctrl}
	mock.recorder = &MockIFileTransferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFileTransfer) EXPECT() *MockIFileTransferMockRecorder {
	return m.recorder
}

// GetFile mocks base method.
func (m *MockIFileTransfer) GetFile(ctx context.Context, archiveID primitive.ID, contentType archive.ContentType) (archive.Data, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", ctx, archiveID, contentType)
	ret0, _ := ret[0].(archive.Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockIFileTransferMockRecorder) GetFile(ctx, archiveID, contentType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockIFileTransfer)(nil).GetFile), ctx, archiveID, contentType)
}

// Save mocks base method.
func (m *MockIFileTransfer) Save(ctx context.Context, archiveID primitive.ID, contentType archive.ContentType, data archive.Data) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, archiveID, contentType, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIFileTransferMockRecorder) Save(ctx, archiveID, contentType, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIFileTransfer)(nil).Save), ctx, archiveID, contentType, data)
}
