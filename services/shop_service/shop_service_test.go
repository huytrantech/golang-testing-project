package shop_service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"golang-testing-project/models"
	"testing"
)

type UnitTestGetShop struct {
	ShopId           int
	Expected         models.GetShopResponse
	ExpectedErrorStr string
	Name             string
}

func TestGetShop(t *testing.T) {
	shopService := NewShopService()
	unitTest := []UnitTestGetShop{
		{
			ShopId: 1,
			Name:   "unit_test_1",
			Expected: models.GetShopResponse{
				ShopId:   1,
				ShopName: "ShopName"},
			ExpectedErrorStr: ""},
		{
			ShopId: 2,
			Name:   "unit_test_2",
			Expected: models.GetShopResponse{
				ShopId:   1,
				ShopName: "ShopName"},
			ExpectedErrorStr: "Shop not found"},
	}
	for _, ut := range unitTest {
		response, err := shopService.GetShop(context.Background(), ut.ShopId)
		if err != nil {
			assert.EqualErrorf(t, err, ut.ExpectedErrorStr, ut.Name)
			continue
		}
		assert.Equalf(t, ut.Expected, response, ut.Name)
	}
}
