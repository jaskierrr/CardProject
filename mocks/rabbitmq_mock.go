// Code generated by MockGen. DO NOT EDIT.
// Source: ./rabbitmq.go

// Package mock is a generated GoMock package.
package mock

import (
	config "card-project/config"
	models "card-project/models"
	rabbitmq "card-project/repositories/rabbitmq"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgx/v5/pgconn"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

// MockuserRepo is a mock of userRepo interface.
type MockuserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockuserRepoMockRecorder
}

// MockuserRepoMockRecorder is the mock recorder for MockuserRepo.
type MockuserRepoMockRecorder struct {
	mock *MockuserRepo
}

// NewMockuserRepo creates a new mock instance.
func NewMockuserRepo(ctrl *gomock.Controller) *MockuserRepo {
	mock := &MockuserRepo{ctrl: ctrl}
	mock.recorder = &MockuserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockuserRepo) EXPECT() *MockuserRepoMockRecorder {
	return m.recorder
}

// DeleteUserID mocks base method.
func (m *MockuserRepo) DeleteUserID(ctx context.Context, id int) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserID", ctx, id)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserID indicates an expected call of DeleteUserID.
func (mr *MockuserRepoMockRecorder) DeleteUserID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserID", reflect.TypeOf((*MockuserRepo)(nil).DeleteUserID), ctx, id)
}

// PostUser mocks base method.
func (m *MockuserRepo) PostUser(ctx context.Context, userData models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostUser", ctx, userData)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostUser indicates an expected call of PostUser.
func (mr *MockuserRepoMockRecorder) PostUser(ctx, userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostUser", reflect.TypeOf((*MockuserRepo)(nil).PostUser), ctx, userData)
}

// MockcardRepo is a mock of cardRepo interface.
type MockcardRepo struct {
	ctrl     *gomock.Controller
	recorder *MockcardRepoMockRecorder
}

// MockcardRepoMockRecorder is the mock recorder for MockcardRepo.
type MockcardRepoMockRecorder struct {
	mock *MockcardRepo
}

// NewMockcardRepo creates a new mock instance.
func NewMockcardRepo(ctrl *gomock.Controller) *MockcardRepo {
	mock := &MockcardRepo{ctrl: ctrl}
	mock.recorder = &MockcardRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcardRepo) EXPECT() *MockcardRepoMockRecorder {
	return m.recorder
}

// DeleteCardID mocks base method.
func (m *MockcardRepo) DeleteCardID(ctx context.Context, id int) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCardID", ctx, id)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCardID indicates an expected call of DeleteCardID.
func (mr *MockcardRepoMockRecorder) DeleteCardID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCardID", reflect.TypeOf((*MockcardRepo)(nil).DeleteCardID), ctx, id)
}

// PostCard mocks base method.
func (m *MockcardRepo) PostCard(ctx context.Context, cardData models.Card) (models.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostCard", ctx, cardData)
	ret0, _ := ret[0].(models.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostCard indicates an expected call of PostCard.
func (mr *MockcardRepoMockRecorder) PostCard(ctx, cardData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostCard", reflect.TypeOf((*MockcardRepo)(nil).PostCard), ctx, cardData)
}

// MockRabbitMQ is a mock of RabbitMQ interface.
type MockRabbitMQ struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitMQMockRecorder
}

// MockRabbitMQMockRecorder is the mock recorder for MockRabbitMQ.
type MockRabbitMQMockRecorder struct {
	mock *MockRabbitMQ
}

// NewMockRabbitMQ creates a new mock instance.
func NewMockRabbitMQ(ctrl *gomock.Controller) *MockRabbitMQ {
	mock := &MockRabbitMQ{ctrl: ctrl}
	mock.recorder = &MockRabbitMQMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitMQ) EXPECT() *MockRabbitMQMockRecorder {
	return m.recorder
}

// NewConn mocks base method.
func (m *MockRabbitMQ) NewConn(userRepo MockuserRepo, cardRepo MockcardRepo, connConfigString string, config config.Config) rabbitmq.RabbitMQ {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewConn", userRepo, cardRepo, connConfigString, config)
	ret0, _ := ret[0].(rabbitmq.RabbitMQ)
	return ret0
}

// NewConn indicates an expected call of NewConn.
func (mr *MockRabbitMQMockRecorder) NewConn(userRepo, cardRepo, connConfigString, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConn", reflect.TypeOf((*MockRabbitMQ)(nil).NewConn), userRepo, cardRepo, connConfigString, config)
}

// NewConsumer mocks base method.
func (m *MockRabbitMQ) NewConsumer(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewConsumer", ctx)
}

// NewConsumer indicates an expected call of NewConsumer.
func (mr *MockRabbitMQMockRecorder) NewConsumer(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConsumer", reflect.TypeOf((*MockRabbitMQ)(nil).NewConsumer), ctx)
}

// ProduceDeleteCard mocks base method.
func (m *MockRabbitMQ) ProduceDeleteCard(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceDeleteCard", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProduceDeleteCard indicates an expected call of ProduceDeleteCard.
func (mr *MockRabbitMQMockRecorder) ProduceDeleteCard(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceDeleteCard", reflect.TypeOf((*MockRabbitMQ)(nil).ProduceDeleteCard), ctx, id)
}

// ProduceDeleteUser mocks base method.
func (m *MockRabbitMQ) ProduceDeleteUser(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceDeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProduceDeleteUser indicates an expected call of ProduceDeleteUser.
func (mr *MockRabbitMQMockRecorder) ProduceDeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceDeleteUser", reflect.TypeOf((*MockRabbitMQ)(nil).ProduceDeleteUser), ctx, id)
}

// ProducePostCard mocks base method.
func (m *MockRabbitMQ) ProducePostCard(ctx context.Context, cardData models.Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProducePostCard", ctx, cardData)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProducePostCard indicates an expected call of ProducePostCard.
func (mr *MockRabbitMQMockRecorder) ProducePostCard(ctx, cardData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProducePostCard", reflect.TypeOf((*MockRabbitMQ)(nil).ProducePostCard), ctx, cardData)
}

// ProducePostUser mocks base method.
func (m *MockRabbitMQ) ProducePostUser(ctx context.Context, userData models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProducePostUser", ctx, userData)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProducePostUser indicates an expected call of ProducePostUser.
func (mr *MockRabbitMQMockRecorder) ProducePostUser(ctx, userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProducePostUser", reflect.TypeOf((*MockRabbitMQ)(nil).ProducePostUser), ctx, userData)
}

// consumeCardDelete mocks base method.
func (m *MockRabbitMQ) consumeCardDelete(ctx context.Context, msg amqp091.Delivery) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "consumeCardDelete", ctx, msg)
}

// consumeCardDelete indicates an expected call of consumeCardDelete.
func (mr *MockRabbitMQMockRecorder) consumeCardDelete(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "consumeCardDelete", reflect.TypeOf((*MockRabbitMQ)(nil).consumeCardDelete), ctx, msg)
}

// consumeCardPost mocks base method.
func (m *MockRabbitMQ) consumeCardPost(ctx context.Context, msg amqp091.Delivery) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "consumeCardPost", ctx, msg)
}

// consumeCardPost indicates an expected call of consumeCardPost.
func (mr *MockRabbitMQMockRecorder) consumeCardPost(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "consumeCardPost", reflect.TypeOf((*MockRabbitMQ)(nil).consumeCardPost), ctx, msg)
}

// consumeUserDelete mocks base method.
func (m *MockRabbitMQ) consumeUserDelete(ctx context.Context, msg amqp091.Delivery) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "consumeUserDelete", ctx, msg)
}

// consumeUserDelete indicates an expected call of consumeUserDelete.
func (mr *MockRabbitMQMockRecorder) consumeUserDelete(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "consumeUserDelete", reflect.TypeOf((*MockRabbitMQ)(nil).consumeUserDelete), ctx, msg)
}

// consumeUserPost mocks base method.
func (m *MockRabbitMQ) consumeUserPost(ctx context.Context, msg amqp091.Delivery) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "consumeUserPost", ctx, msg)
}

// consumeUserPost indicates an expected call of consumeUserPost.
func (mr *MockRabbitMQMockRecorder) consumeUserPost(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "consumeUserPost", reflect.TypeOf((*MockRabbitMQ)(nil).consumeUserPost), ctx, msg)
}
