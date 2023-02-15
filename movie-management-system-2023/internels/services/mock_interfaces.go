// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	gofr "developer.zopsmart.com/go/gofr/pkg/gofr"
	models "github.com/go-training/movie-management-system-2023/internels/models"
	gomock "github.com/golang/mock/gomock"
)

// MockMovieManager is a mock of MovieManager interface.
type MockMovieManager struct {
	ctrl     *gomock.Controller
	recorder *MockMovieManagerMockRecorder
}

// MockMovieManagerMockRecorder is the mock recorder for MockMovieManager.
type MockMovieManagerMockRecorder struct {
	mock *MockMovieManager
}

// NewMockMovieManager creates a new mock instance.
func NewMockMovieManager(ctrl *gomock.Controller) *MockMovieManager {
	mock := &MockMovieManager{ctrl: ctrl}
	mock.recorder = &MockMovieManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieManager) EXPECT() *MockMovieManagerMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockMovieManager) Delete(ctx *gofr.Context, movieID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, movieID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockMovieManagerMockRecorder) Delete(ctx, movieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMovieManager)(nil).Delete), ctx, movieID)
}

// Get mocks base method.
func (m *MockMovieManager) Get(ctx *gofr.Context, movieID int) (models.Movie, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, movieID)
	ret0, _ := ret[0].(models.Movie)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockMovieManagerMockRecorder) Get(ctx, movieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMovieManager)(nil).Get), ctx, movieID)
}

// GetAll mocks base method.
func (m *MockMovieManager) GetAll(ctx *gofr.Context) ([]models.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]models.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockMovieManagerMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockMovieManager)(nil).GetAll), ctx)
}

// Post mocks base method.
func (m *MockMovieManager) Post(ctx *gofr.Context, mov *models.Movie) (models.Movie, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", ctx, mov)
	ret0, _ := ret[0].(models.Movie)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Post indicates an expected call of Post.
func (mr *MockMovieManagerMockRecorder) Post(ctx, mov interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockMovieManager)(nil).Post), ctx, mov)
}

// Update mocks base method.
func (m *MockMovieManager) Update(ctx *gofr.Context, movieID int, upMovie *models.UpdateMovie) (models.Movie, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, movieID, upMovie)
	ret0, _ := ret[0].(models.Movie)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockMovieManagerMockRecorder) Update(ctx, movieID, upMovie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMovieManager)(nil).Update), ctx, movieID, upMovie)
}