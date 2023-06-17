// Code generated by MockGen. DO NOT EDIT.
// Source: sd/internal/domain/animal (interfaces: IAnimalRepo)

// Package animal is a generated GoMock package.
package animal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAnimalRepo is a mock of IAnimalRepo interface.
type MockIAnimalRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIAnimalRepoMockRecorder
}

// MockIAnimalRepoMockRecorder is the mock recorder for MockIAnimalRepo.
type MockIAnimalRepoMockRecorder struct {
	mock *MockIAnimalRepo
}

// NewMockIAnimalRepo creates a new mock instance.
func NewMockIAnimalRepo(ctrl *gomock.Controller) *MockIAnimalRepo {
	mock := &MockIAnimalRepo{ctrl: ctrl}
	mock.recorder = &MockIAnimalRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAnimalRepo) EXPECT() *MockIAnimalRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIAnimalRepo) Create(arg0 context.Context, arg1 *CreateAnimal) (*Animal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*Animal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIAnimalRepoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIAnimalRepo)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIAnimalRepo) Delete(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIAnimalRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIAnimalRepo)(nil).Delete), arg0, arg1)
}

// GetAll mocks base method.
func (m *MockIAnimalRepo) GetAll(arg0 context.Context) (Animals, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].(Animals)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIAnimalRepoMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIAnimalRepo)(nil).GetAll), arg0)
}

// GetById mocks base method.
func (m *MockIAnimalRepo) GetById(arg0 context.Context, arg1 int32) (*Animal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*Animal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIAnimalRepoMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIAnimalRepo)(nil).GetById), arg0, arg1)
}

// Update mocks base method.
func (m *MockIAnimalRepo) Update(arg0 context.Context, arg1 int32, arg2 *UpdateAnimal) (*Animal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*Animal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIAnimalRepoMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIAnimalRepo)(nil).Update), arg0, arg1, arg2)
}
