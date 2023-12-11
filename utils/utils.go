package utils

import "golang-testing-project/models"

func CalculateTotalAmount(order models.OrderEntity) (totalAmount float64) {

	totalAmount = order.SubAmount + order.DeclarePriceFee + order.ShippingFee
	if order.VoucherType == 1 { //1: giam theo gia tri voucher
		totalAmount -= order.VoucherValue
	} else if order.VoucherType == 2 { //giam theo gia tri % don hang
		totalAmount -= totalAmount * order.VoucherValue
	}
	if totalAmount < 0 {
		return 0
	}

	return
}

func GetRoles(status int) (resp models.OrderRoles) {
	if status == 1 {
		resp.CanConfirm = true
		resp.CanCancel = true
	} else if status == 2 {
		resp.CanCancel = true
	}
	return
}
