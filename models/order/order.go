package order

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"

	// mysqlDb "github.com/Sahil-Mandaliya/Order/backend/pkg/mysql"
	mysqlDb "github.com/Sahil-Mandaliya/OrderService/pkg/mysql"
)

type Order struct {
	ID        uint64    `gorm:"id" json:"id"`
	UserID    uint64    `gorm:"user_id" json:"user_id"`
	Price     float64   `gorm:"price" json:"price"`
	OrderDate time.Time `gorm:"order_date" json:"order_date"`
	IsDeleted uint8     `gorm:"is_deleted" json:"is_deleted"`
}

// Create order
func CreateOrder(ctx context.Context, order *Order) (*Order, error) {
	db := mysqlDb.DBConnection()
	tx := db.Create(order)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return order, nil
}

// Get all orders
func GetOrders(ctx context.Context) ([]*Order, error) {
	db := mysqlDb.DBConnection()
	var orders []*Order
	tx := db.Model(&Order{}).Where("is_deleted=0").Find(&orders)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return orders, nil
}

// Get all orders
func GetOrderByOrderId(ctx context.Context, orderId uint64) (*Order, error) {
	db := mysqlDb.DBConnection()
	order := &Order{}
	tx := db.Model(&Order{}).Where("id=?", orderId).Last(order)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return order, nil
}

// Update Order
func UpdateOrder(ctx context.Context, order *Order) (*Order, error) {
	db := mysqlDb.DBConnection()
	tx := db.UpdateColumns(order).Where("id=?", order.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return order, nil
}

// Delete Order
func DeleteOrder(ctx context.Context, orderID uint64) error {
	db := mysqlDb.DBConnection()
	columns := map[string]interface{}{
		"is_deleted": 1,
	}
	tx := db.Model(&Order{}).Where("id=?", orderID).Updates(columns)
	return tx.Error
}
