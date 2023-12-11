package utils

import (
	"github.com/stretchr/testify/assert"
	"golang-testing-project/models"
	"testing"
)

type UnitTestCalculateTotalAmount struct {
	Order    models.OrderEntity
	Expected float64
	Name     string
}

func TestCalculateTotalAmount(t *testing.T) {

	unitTest := []UnitTestCalculateTotalAmount{
		{
			Order: models.OrderEntity{
				OrderCode:       "1",
				ShopId:          1,
				DeclarePriceFee: 10000,
				VoucherValue:    20000,
				SubAmount:       200000,
				ShippingFee:     15000,
				VoucherType:     1,
			},
			Expected: 205000,
			Name:     "unit_test_1",
		},
		{
			Order: models.OrderEntity{
				OrderCode:       "1",
				ShopId:          1,
				DeclarePriceFee: 0,
				VoucherValue:    20000,
				SubAmount:       200000,
				ShippingFee:     15000,
				VoucherType:     1,
			},
			Expected: 195000,
			Name:     "unit_test_2",
		},
		{
			Order: models.OrderEntity{
				OrderCode:       "1",
				ShopId:          1,
				DeclarePriceFee: 0,
				VoucherValue:    20000,
				SubAmount:       200000,
				ShippingFee:     0,
				VoucherType:     1,
			},
			Expected: 180000,
			Name:     "unit_test_3",
		},
		{
			Order: models.OrderEntity{
				OrderCode:       "1",
				ShopId:          1,
				DeclarePriceFee: 0,
				VoucherValue:    20000,
				SubAmount:       15000,
				ShippingFee:     0,
				VoucherType:     1,
			},
			Expected: 0,
			Name:     "unit_test_4",
		},
		{
			Order: models.OrderEntity{
				OrderCode:       "1",
				ShopId:          1,
				DeclarePriceFee: 0,
				VoucherValue:    15000,
				SubAmount:       15000,
				ShippingFee:     0,
				VoucherType:     1,
			},
			Expected: 0,
			Name:     "unit_test_5",
		},
		{
			Order: models.OrderEntity{
				OrderCode:       "1",
				ShopId:          1,
				DeclarePriceFee: 0,
				VoucherValue:    0.2,
				SubAmount:       200000,
				ShippingFee:     0,
				VoucherType:     2,
			},
			Expected: 160000,
			Name:     "unit_test_6",
		},
		{
			Order: models.OrderEntity{
				OrderCode:       "1",
				ShopId:          1,
				DeclarePriceFee: 0,
				VoucherValue:    0,
				SubAmount:       200000,
				ShippingFee:     0,
				VoucherType:     0,
			},
			Expected: 200000,
			Name:     "unit_test_7",
		},
	}

	for _, ut := range unitTest {
		amount := CalculateTotalAmount(ut.Order)
		assert.Equalf(t, ut.Expected, amount, ut.Name)
	}
}
