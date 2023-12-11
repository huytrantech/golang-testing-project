//mockgen --source=./repository/order_repo.go --destination=./mocking/order_repo_mockgen.go --package=mocking

package repository

import (
	"context"
	"errors"
	"golang-testing-project/models"
)

type IOrderRepository interface {
	FindOrder(ctx context.Context, orderCode string) (*models.OrderEntity, error)
}

type orderRepository struct {
}

func (o *orderRepository) FindOrder(ctx context.Context, orderCode string) (order *models.OrderEntity, err error) {
	if orderCode == "123456789" {
		order = &models.OrderEntity{
			OrderCode:       "123456789",
			ShopId:          1,
			DeclarePriceFee: 10000,
			VoucherValue:    20000,
			SubAmount:       200000,
			ShippingFee:     15000,
			VoucherType:     1,
			Status:          1,
		}
	} else if orderCode == "111" {
		return
	} else {
		err = errors.New("internal error")
	}
	return
}

func NewOrderRepository() IOrderRepository {
	return &orderRepository{}
}
