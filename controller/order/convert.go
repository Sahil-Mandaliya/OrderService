package order

import (

	orderModel "github.com/Sahil-Mandaliya/OrderService/models/order"
	"github.com/Sahil-Mandaliya/OrderService/pkg/util"
	orderpb "github.com/Sahil-Mandaliya/OrderService/proto/order"
)

func OrdersModelToPb(orders []*orderModel.Order) []*orderpb.Order {
	ordersPb := make([]*orderpb.Order, 0)
	for _, order := range orders {
		ordersPb = append(ordersPb, OrderModelToPb(order))
	}
	return ordersPb
}

func OrderModelToPb(order *orderModel.Order) *orderpb.Order {
	orderPb := &orderpb.Order{
		OrderId:   order.ID,
		UserId:    order.UserID,
		Price:     float32(order.Price),
		OrderDate: util.FormatTimeToStringDDMMYYYY(order.OrderDate),
	}
	return orderPb
}
