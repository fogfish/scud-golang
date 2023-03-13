// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fogfish/blueprint-serverless-golang/http (interfaces: PetFetcher,PetCreator)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	core "github.com/fogfish/blueprint-serverless-golang"
	gomock "github.com/golang/mock/gomock"
)

// MockPetFetcher is a mock of PetFetcher interface.
type MockPetFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockPetFetcherMockRecorder
}

// MockPetFetcherMockRecorder is the mock recorder for MockPetFetcher.
type MockPetFetcherMockRecorder struct {
	mock *MockPetFetcher
}

// NewMockPetFetcher creates a new mock instance.
func NewMockPetFetcher(ctrl *gomock.Controller) *MockPetFetcher {
	mock := &MockPetFetcher{ctrl: ctrl}
	mock.recorder = &MockPetFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPetFetcher) EXPECT() *MockPetFetcherMockRecorder {
	return m.recorder
}

// LookupPet mocks base method.
func (m *MockPetFetcher) LookupPet(arg0 context.Context, arg1 string) (core.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupPet", arg0, arg1)
	ret0, _ := ret[0].(core.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupPet indicates an expected call of LookupPet.
func (mr *MockPetFetcherMockRecorder) LookupPet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupPet", reflect.TypeOf((*MockPetFetcher)(nil).LookupPet), arg0, arg1)
}

// LookupPetsAfterKey mocks base method.
func (m *MockPetFetcher) LookupPetsAfterKey(arg0 context.Context, arg1 string, arg2 int) ([]core.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupPetsAfterKey", arg0, arg1, arg2)
	ret0, _ := ret[0].([]core.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupPetsAfterKey indicates an expected call of LookupPetsAfterKey.
func (mr *MockPetFetcherMockRecorder) LookupPetsAfterKey(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupPetsAfterKey", reflect.TypeOf((*MockPetFetcher)(nil).LookupPetsAfterKey), arg0, arg1, arg2)
}

// MockPetCreator is a mock of PetCreator interface.
type MockPetCreator struct {
	ctrl     *gomock.Controller
	recorder *MockPetCreatorMockRecorder
}

// MockPetCreatorMockRecorder is the mock recorder for MockPetCreator.
type MockPetCreatorMockRecorder struct {
	mock *MockPetCreator
}

// NewMockPetCreator creates a new mock instance.
func NewMockPetCreator(ctrl *gomock.Controller) *MockPetCreator {
	mock := &MockPetCreator{ctrl: ctrl}
	mock.recorder = &MockPetCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPetCreator) EXPECT() *MockPetCreatorMockRecorder {
	return m.recorder
}

// CreatePet mocks base method.
func (m *MockPetCreator) CreatePet(arg0 context.Context, arg1 core.Pet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePet indicates an expected call of CreatePet.
func (mr *MockPetCreatorMockRecorder) CreatePet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePet", reflect.TypeOf((*MockPetCreator)(nil).CreatePet), arg0, arg1)
}