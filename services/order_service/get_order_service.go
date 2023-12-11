package order_service

import (
	"context"
	"errors"
	"golang-testing-project/models"
	"golang-testing-project/repository"
	"golang-testing-project/services/shop_service"
	"golang-testing-project/utils"
)

type IGetOrderService interface {
	GetOrder(ctx context.Context, orderCode string) (resp models.GetOrderResponse, err error)
}

type getOrderService struct {
	IShopService     shop_service.IShopService
	IOrderRepository repository.IOrderRepository
}

func (g *getOrderService) GetOrder(ctx context.Context, orderCode string) (resp models.GetOrderResponse, err error) {
	orderData, err := g.IOrderRepository.FindOrder(ctx, orderCode)
	if err != nil {
		return
	}

	if orderData == nil {
		err = errors.New("order not found")
		return
	}

	shopData, err := g.IShopService.GetShop(ctx, orderData.ShopId)
	if err != nil {
		return
	}

	resp.ShopId = orderData.ShopId
	resp.ShopName = shopData.ShopName
	resp.OrderCode = orderData.OrderCode
	resp.VoucherValue = orderData.VoucherValue
	resp.VoucherType = orderData.VoucherType
	resp.ShippingFee = orderData.ShippingFee
	resp.DeclarePriceFee = orderData.DeclarePriceFee
	resp.SubAmount = orderData.SubAmount
	resp.TotalAmount = utils.CalculateTotalAmount(*orderData)

	return
}

func NewGetOrderService(IShopService shop_service.IShopService, IOrderRepository repository.IOrderRepository) IGetOrderService {
	return &getOrderService{IShopService: IShopService, IOrderRepository: IOrderRepository}
}
