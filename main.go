package main

import (

	// mysqlDb "github.com/Sahil-Mandaliya/Order/backend/pkg/mysql"
	// routes "github.com/Sahil-Mandaliya/Order/backend/routes"

	"log"
	"net"

	routes "github.com/Sahil-Mandaliya/OrderService/routes"
	"google.golang.org/grpc"

	orderSvc "github.com/Sahil-Mandaliya/OrderService/Service/order"
	mysqlDb "github.com/Sahil-Mandaliya/OrderService/pkg/mysql"
	orderpb "github.com/Sahil-Mandaliya/OrderService/proto/order"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := mysqlDb.InitDB()
	if db.Error != nil {
		log.Fatalln(db.Error)
	}
	connection := mysqlDb.DBConnection()
	condb := connection.DB()
	defer condb.Close()

	startServices()

	routes.Routers()
}

func startServices() {
	listener, err := net.Listen("tcp", ":6000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	orderpb.RegisterOrderServiceServer(s, &orderSvc.OrderService{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Service listening at 6000")
	// reflection.Register(s)
}
