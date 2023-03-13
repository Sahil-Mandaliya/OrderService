package order

import (
	"context"

	orderpb "github.com/Sahil-Mandaliya/OrderService/proto/order"
)

type OrderService struct {
	orderpb.UnimplementedOrderServiceServer
}

func (rpc OrderService) GetAllOrders(ctx context.Context, req *orderpb.GetAllOrdersRequest) (*orderpb.GetAllOrdersResponse, error) {
	orders := make([]*orderpb.Order, 0)
	res := &orderpb.GetAllOrdersResponse{
		Orders: &orderpb.Orders{
			Order: orders,
		},
	}
	return res, nil
}

func (rpc OrderService) GetOrderById(ctx context.Context, req *orderpb.GetOrderByIdRequest) (*orderpb.GetOrderByIdResponse, error) {
	res := &orderpb.GetOrderByIdResponse{
		Order: &orderpb.Order{},
	}
	return res, nil
}
