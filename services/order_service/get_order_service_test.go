package order_service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-testing-project/mocking"
	"golang-testing-project/models"
	"golang-testing-project/utils"
	"testing"
)

type UnitTestGetOrder struct {
	Name          string
	Expected      models.GetOrderResponse
	CtrlList      func()
	ExpectedError error
	OrderCode     string
}

func TestGetOrderService(t *testing.T) {
	expectedOrder := models.OrderEntity{
		OrderCode:       "123456789",
		ShopId:          1,
		DeclarePriceFee: 10000,
		VoucherValue:    20000,
		SubAmount:       200000,
		ShippingFee:     15000,
		VoucherType:     1,
		Status:          1,
	}
	ctrl := gomock.NewController(t)
	orderRepo := mocking.NewMockIOrderRepository(ctrl)
	NewShopService := mocking.NewMockIShopService(ctrl)
	unitTests := []UnitTestGetOrder{
		{
			Name:      "unit_test_1",
			OrderCode: "123456789",
			CtrlList: func() {
				orderRepo.EXPECT().FindOrder(context.Background(), expectedOrder.OrderCode).Times(1).Return(&expectedOrder, nil)
				NewShopService.EXPECT().GetShop(context.Background(), expectedOrder.ShopId).Times(1).Return(models.GetShopResponse{
					ShopId:   1,
					ShopName: "ShopName1",
				}, nil)
			},
			Expected: models.GetOrderResponse{
				OrderCode:       "123456789",
				DeclarePriceFee: 10000,
				VoucherValue:    20000,
				SubAmount:       200000,
				ShippingFee:     15000,
				VoucherType:     1,
				ShopId:          1,
				ShopName:        "ShopName1",
				TotalAmount:     utils.CalculateTotalAmount(expectedOrder),
			},
		},
		{
			Name:      "unit_test_2",
			OrderCode: "123456789",
			CtrlList: func() {
				orderRepo.EXPECT().FindOrder(context.Background(), expectedOrder.OrderCode).Times(1).Return(nil, nil)
			},
			Expected: models.GetOrderResponse{
				OrderCode:       "123456789",
				DeclarePriceFee: 10000,
				VoucherValue:    20000,
				SubAmount:       200000,
				ShippingFee:     15000,
				VoucherType:     1,
				ShopId:          1,
				ShopName:        "ShopName1",
				TotalAmount:     utils.CalculateTotalAmount(expectedOrder),
			},
			ExpectedError: errors.New("order not found"),
		},
		{
			Name:      "unit_test_3",
			OrderCode: "123456789",
			CtrlList: func() {
				orderRepo.EXPECT().FindOrder(context.Background(), expectedOrder.OrderCode).Times(1).Return(nil, errors.New("panic"))
			},
			Expected: models.GetOrderResponse{
				OrderCode:       "123456789",
				DeclarePriceFee: 10000,
				VoucherValue:    20000,
				SubAmount:       200000,
				ShippingFee:     15000,
				VoucherType:     1,
				ShopId:          1,
				ShopName:        "ShopName1",
				TotalAmount:     utils.CalculateTotalAmount(expectedOrder),
			},
			ExpectedError: errors.New("panic"),
		},
		{
			Name:      "unit_test_4",
			OrderCode: "123456789",
			CtrlList: func() {
				orderRepo.EXPECT().FindOrder(context.Background(), expectedOrder.OrderCode).Times(1).Return(&expectedOrder, nil)
				NewShopService.EXPECT().GetShop(context.Background(), expectedOrder.ShopId).Times(1).Return(models.GetShopResponse{}, errors.New("shop panic"))
			},
			Expected: models.GetOrderResponse{
				OrderCode:       "123456789",
				DeclarePriceFee: 10000,
				VoucherValue:    20000,
				SubAmount:       200000,
				ShippingFee:     15000,
				VoucherType:     1,
				ShopId:          1,
				ShopName:        "ShopName1",
				TotalAmount:     utils.CalculateTotalAmount(expectedOrder),
			},
			ExpectedError: errors.New("shop panic"),
		},
	}
	for _, ut := range unitTests {
		ut.CtrlList()
		NewGetOrderService := NewGetOrderService(NewShopService, orderRepo)
		resp, err := NewGetOrderService.GetOrder(context.Background(), ut.OrderCode)
		if err != nil {
			assert.EqualError(t, ut.ExpectedError, err.Error(), ut.Name)
		} else {
			assert.Equalf(t, ut.Expected, resp, ut.Name)
		}
	}
}
