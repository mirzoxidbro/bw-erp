package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type CreateOrderModel struct {
	CompanyID   string `json:"company_id", binding:"required"`
	Phone       string `json:"phone", binding:"required"`
	Count       int    `json:"count", binding:"required"`
	Slug        string `json:"slug", binding:"required"`
	Description string `json:"description"`
}

type OrderList struct {
	ID     int        `json:"id"`
	Slug   string     `json:"slug"`
	Status NullString `json:"status"`
}

type Order struct {
	ID          int         `json:"id"`
	CompanyID   string      `json:"company_id"`
	Phone       string      `json:"phone"`
	Count       int         `json:"count"`
	Slug        string      `json:"slug"`
	Description string      `json:"description"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Latitute    *float64    `json:"latitute"`
	Longitude   *float64    `json:"longitude"`
	OrderItems  []OrderItem `json:"order_items"`
}

type UpdateOrderRequest struct {
	ID          int    `json:"id", binding:"required"`
	Slug        string `json:"slug, omitempty"`
	Status      int16  `json:"status, omitempty"`
	Phone       string `json:"phone, omitempty"`
	Description string `json:"description"`
}

type NullFloat struct {
	sql.NullFloat64
}

func (ns NullFloat) MarshalJSONFloat() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Float64)
	}
	return json.Marshal(nil)
}