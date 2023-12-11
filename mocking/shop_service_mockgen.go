// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/shop_service/shop_service.go

// Package mocking is a generated GoMock package.
package mocking

import (
	context "context"
	models "golang-testing-project/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIShopService is a mock of IShopService interface.
type MockIShopService struct {
	ctrl     *gomock.Controller
	recorder *MockIShopServiceMockRecorder
}

// MockIShopServiceMockRecorder is the mock recorder for MockIShopService.
type MockIShopServiceMockRecorder struct {
	mock *MockIShopService
}

// NewMockIShopService creates a new mock instance.
func NewMockIShopService(ctrl *gomock.Controller) *MockIShopService {
	mock := &MockIShopService{ctrl: ctrl}
	mock.recorder = &MockIShopServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIShopService) EXPECT() *MockIShopServiceMockRecorder {
	return m.recorder
}

// GetShop mocks base method.
func (m *MockIShopService) GetShop(ctx context.Context, shopId int) (models.GetShopResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShop", ctx, shopId)
	ret0, _ := ret[0].(models.GetShopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShop indicates an expected call of GetShop.
func (mr *MockIShopServiceMockRecorder) GetShop(ctx, shopId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShop", reflect.TypeOf((*MockIShopService)(nil).GetShop), ctx, shopId)
}