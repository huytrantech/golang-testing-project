package models

type OrderEntity struct {
	OrderCode       string
	ShopId          int
	DeclarePriceFee float64
	VoucherValue    float64
	SubAmount       float64
	ShippingFee     float64
	VoucherType     int
	Status          int
}

type GetOrderResponse struct {
	OrderCode       string
	DeclarePriceFee float64
	VoucherValue    float64
	SubAmount       float64
	ShippingFee     float64
	VoucherType     int
	ShopId          int
	ShopName        string
	TotalAmount     float64
}

type OrderRoles struct {
	CanCancel  bool
	CanConfirm bool
}
