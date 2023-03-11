package routes

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	orderCtr "github.com/Sahil-Mandaliya/OrderService/controller/order"
	userCtr "github.com/Sahil-Mandaliya/OrderService/controller/user"
)

func Routers() {
	router := mux.NewRouter()
	router.HandleFunc("/users",
		userCtr.GetUsers).Methods("GET")
	router.HandleFunc("/user",
		userCtr.GetUserById).Methods("GET")
	router.HandleFunc("/orders",
		orderCtr.GetOrders).Methods("GET")
	router.HandleFunc("/users",
		userCtr.CreateUser).Methods("POST")
	router.HandleFunc("/order",
		orderCtr.CreateOrder).Methods("POST")
	router.HandleFunc("/order",
		orderCtr.GetOrder).Methods("GET")
	// ?order_id={id}
	router.HandleFunc("/order/update",
		orderCtr.UpdateOrder).Methods("PUT")
	router.HandleFunc("/order/delete",
		orderCtr.DeleteOrder).Methods("DELETE")

	// router.HandleFunc("/users/{id}",
	// 	userModel.GetUser).Methods("GET")
	// router.HandleFunc("/users/{id}",
	// 	userModel.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users/{id}",
	// 	userModel.DeleteUser).Methods("DELETE")
	http.ListenAndServe(":9080",
		&CORSRouterDecorator{router})
}

/***************************************************/

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Accept-Language,"+
				" Content-Type, YourOwnHeader")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}
