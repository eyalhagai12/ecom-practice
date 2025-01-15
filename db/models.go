// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

func (e *OrderStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OrderStatus(s)
	case string:
		*e = OrderStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for OrderStatus: %T", src)
	}
	return nil
}

type NullOrderStatus struct {
	OrderStatus OrderStatus `json:"orderStatus"`
	Valid       bool        `json:"valid"` // Valid is true if OrderStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrderStatus) Scan(value interface{}) error {
	if value == nil {
		ns.OrderStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OrderStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrderStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OrderStatus), nil
}

type Inventory struct {
	ID         uuid.UUID        `json:"id"`
	ProductID  uuid.UUID        `json:"productId"`
	Quantity   int32            `json:"quantity"`
	LocationID int64            `json:"locationId"`
	CreatedAt  pgtype.Timestamp `json:"createdAt"`
	UpdatedAt  pgtype.Timestamp `json:"updatedAt"`
	DeletedAt  pgtype.Timestamp `json:"deletedAt"`
}

type Location struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Order struct {
	ID         uuid.UUID        `json:"id"`
	CreatedAt  pgtype.Timestamp `json:"createdAt"`
	UpdatedAt  pgtype.Timestamp `json:"updatedAt"`
	Status     OrderStatus      `json:"status"`
	TotalPrice float64          `json:"totalPrice"`
}

type OrderItem struct {
	ID        int64            `json:"id"`
	OrderID   uuid.UUID        `json:"orderId"`
	ProductID uuid.UUID        `json:"productId"`
	Quantity  int32            `json:"quantity"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
}

type Product struct {
	ID        uuid.UUID        `json:"id"`
	Name      string           `json:"name"`
	Price     float64          `json:"price"`
	Quantity  *int32           `json:"quantity"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	DeletedAt pgtype.Timestamp `json:"deletedAt"`
	StoreID   uuid.UUID        `json:"storeId"`
}

type Store struct {
	ID        uuid.UUID        `json:"id"`
	Name      string           `json:"name"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	DeletedAt pgtype.Timestamp `json:"deletedAt"`
}
