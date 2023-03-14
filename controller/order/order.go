package order

import (
	"context"
	"net/http"

	orderModel "github.com/Sahil-Mandaliya/OrderService/models/order"
	jsonEncode "github.com/Sahil-Mandaliya/OrderService/pkg/json"
	"github.com/Sahil-Mandaliya/OrderService/pkg/util"
)

type Response struct {
	Status   bool
	Message  string
	Response interface{}
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	order := &orderModel.Order{}
	jsonEncode.HttpRequestUnmarshal(r, order)
	order, err := orderModel.CreateOrder(ctx, order)
	if err != nil {
		response := &Response{
			Status:  false,
			Message: err.Error(),
		}
		jsonEncode.JsonResposeEncoder(w, r, response)
		return
	}
	response := &Response{
		Status:   true,
		Message:  "success",
		Response: order,
	}
	jsonEncode.JsonResposeEncoder(w, r, response)
}

// func GetOrders(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	orders, err := orderModel.GetOrders(ctx)
// 	if err != nil {
// 		response := &Response{
// 			Status:  false,
// 			Message: err.Error(),
// 		}
// 		jsonEncode.JsonResposeEncoder(w, r, response)
// 		return
// 	}
// 	response := &Response{
// 		Status:   true,
// 		Message:  "success",
// 		Response: orders,
// 	}
// 	jsonEncode.JsonResposeEncoder(w, r, response)
// }

func GetOrders(ctx context.Context) ([]*orderModel.Order, error) {
	orders, err := orderModel.GetOrders(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// func GetOrder(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	w.Header().Set("Content-Type", "application/json")
// 	queryParams := r.URL.Query()
// 	orderId, exists := queryParams["order_id"]
// 	if !exists {
// 		response := &Response{
// 			Status:  false,
// 			Message: "Invalid OrderId",
// 		}
// 		jsonEncode.JsonResposeEncoder(w, r, response)
// 		return
// 	}
// 	orders, err := orderModel.GetOrderByOrderId(ctx, util.ParseUint64(orderId[0]))
// 	if err != nil {
// 		response := &Response{
// 			Status:  false,
// 			Message: err.Error(),
// 		}
// 		jsonEncode.JsonResposeEncoder(w, r, response)
// 		return
// 	}
// 	response := &Response{
// 		Status:   true,
// 		Message:  "success",
// 		Response: orders,
// 	}
// 	jsonEncode.JsonResposeEncoder(w, r, response)
// }

func GetOrder(ctx context.Context, orderId uint64) (*orderModel.Order, error) {
	order, err := orderModel.GetOrderByOrderId(ctx, orderId)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParams := r.URL.Query()
	orderId, exists := queryParams["order_id"]
	if !exists {
		response := &Response{
			Status:  false,
			Message: "Invalid OrderId",
		}
		jsonEncode.JsonResposeEncoder(w, r, response)
		return
	}
	order := &orderModel.Order{}
	jsonEncode.HttpRequestUnmarshal(r, order)
	order.ID = util.ParseUint64(orderId[0])
	order, err := orderModel.UpdateOrder(ctx, order)
	if err != nil {
		response := &Response{
			Status:  false,
			Message: err.Error(),
		}
		jsonEncode.JsonResposeEncoder(w, r, response)
		return
	}
	response := &Response{
		Status:   true,
		Message:  "success",
		Response: order,
	}
	jsonEncode.JsonResposeEncoder(w, r, response)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParams := r.URL.Query()
	orderIds, exists := queryParams["order_id"]
	if !exists {
		response := &Response{
			Status:  false,
			Message: "Invalid OrderId",
		}
		jsonEncode.JsonResposeEncoder(w, r, response)
		return
	}
	orderId := util.ParseUint64(orderIds[0])
	err := orderModel.DeleteOrder(ctx, orderId)
	if err != nil {
		response := &Response{
			Status:  false,
			Message: err.Error(),
		}
		jsonEncode.JsonResposeEncoder(w, r, response)
		return
	}
	response := &Response{
		Status:  true,
		Message: "successfully deleted order",
	}
	jsonEncode.JsonResposeEncoder(w, r, response)
}
