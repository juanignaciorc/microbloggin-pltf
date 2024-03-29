// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/ports/services/users_services.go
//
// Generated by this command:
//
//	mockgen -source=../internal/ports/services/users_services.go -destination=./mock_users_service.go
//

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	domain "github.com/juanignaciorc/microbloggin-pltf/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserService) CreateUser(ctx context.Context, name, mail string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, name, mail)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceMockRecorder) CreateUser(ctx, name, mail any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), ctx, name, mail)
}

// FollowUser mocks base method.
func (m *MockUserService) FollowUser(ctx context.Context, userID, followedID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FollowUser", ctx, userID, followedID)
	ret0, _ := ret[0].(error)
	return ret0
}

// FollowUser indicates an expected call of FollowUser.
func (mr *MockUserServiceMockRecorder) FollowUser(ctx, userID, followedID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FollowUser", reflect.TypeOf((*MockUserService)(nil).FollowUser), ctx, userID, followedID)
}

// GetUser mocks base method.
func (m *MockUserService) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserServiceMockRecorder) GetUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserService)(nil).GetUser), ctx, id)
}

// GetUserTimeline mocks base method.
func (m *MockUserService) GetUserTimeline(ctx context.Context, userID uuid.UUID) ([]domain.Tweet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserTimeline", ctx, userID)
	ret0, _ := ret[0].([]domain.Tweet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserTimeline indicates an expected call of GetUserTimeline.
func (mr *MockUserServiceMockRecorder) GetUserTimeline(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserTimeline", reflect.TypeOf((*MockUserService)(nil).GetUserTimeline), ctx, userID)
}
