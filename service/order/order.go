package order

import (
	"context"

	orderCtr "github.com/Sahil-Mandaliya/OrderService/controller/order"
	orderpb "github.com/Sahil-Mandaliya/OrderService/proto/order"
)

type OrderService struct {
	orderpb.UnimplementedOrderServiceServer
}

func (rpc OrderService) GetAllOrders(ctx context.Context, req *orderpb.GetAllOrdersRequest) (*orderpb.GetAllOrdersResponse, error) {
	orders, err := orderCtr.GetOrders(ctx)
	if err != nil {
		return nil, err
	}
	res := &orderpb.GetAllOrdersResponse{
		Orders: orderCtr.OrdersModelToPb(orders),
	}
	return res, nil
}

func (rpc OrderService) GetOrderById(ctx context.Context, req *orderpb.GetOrderByIdRequest) (*orderpb.GetOrderByIdResponse, error) {
	order, err := orderCtr.GetOrder(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	res := &orderpb.GetOrderByIdResponse{
		Order: orderCtr.OrderModelToPb(order),
	}
	return res, nil
}
