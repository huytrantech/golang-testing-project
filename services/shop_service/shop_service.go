package shop_service

//mockgen --source=./services/shop_service/shop_service.go --destination=./mocking/shop_service_mockgen.go --package=mocking
import (
	"context"
	"errors"
	"golang-testing-project/models"
)

type IShopService interface {
	GetShop(ctx context.Context, shopId int) (resp models.GetShopResponse, err error)
}

type shopService struct {
}

func (s shopService) GetShop(ctx context.Context, shopId int) (resp models.GetShopResponse, err error) {
	if shopId == 1 {
		resp.ShopId = 1
		resp.ShopName = "ShopName"
		return
	}
	err = errors.New("Shop not found")
	return
}

func NewShopService() IShopService {
	return &shopService{}
}
