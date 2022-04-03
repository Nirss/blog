// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nirss/blog/internal/service (interfaces: BlogService)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	post "github.com/Nirss/blog/internal/domain/post"
	gomock "github.com/golang/mock/gomock"
)

// MockBlogService is a mock of BlogService interface.
type MockBlogService struct {
	ctrl     *gomock.Controller
	recorder *MockBlogServiceMockRecorder
}

// MockBlogServiceMockRecorder is the mock recorder for MockBlogService.
type MockBlogServiceMockRecorder struct {
	mock *MockBlogService
}

// NewMockBlogService creates a new mock instance.
func NewMockBlogService(ctrl *gomock.Controller) *MockBlogService {
	mock := &MockBlogService{ctrl: ctrl}
	mock.recorder = &MockBlogServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlogService) EXPECT() *MockBlogServiceMockRecorder {
	return m.recorder
}

// CreateRecord mocks base method.
func (m *MockBlogService) CreateRecord(arg0 context.Context, arg1 post.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecord", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRecord indicates an expected call of CreateRecord.
func (mr *MockBlogServiceMockRecorder) CreateRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecord", reflect.TypeOf((*MockBlogService)(nil).CreateRecord), arg0, arg1)
}

// GetAllRecords mocks base method.
func (m *MockBlogService) GetAllRecords(arg0 context.Context) ([]post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRecords", arg0)
	ret0, _ := ret[0].([]post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRecords indicates an expected call of GetAllRecords.
func (mr *MockBlogServiceMockRecorder) GetAllRecords(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRecords", reflect.TypeOf((*MockBlogService)(nil).GetAllRecords), arg0)
}