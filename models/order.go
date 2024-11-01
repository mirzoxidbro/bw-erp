package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type PaymentType string
type PaymentStatus int8

const (
	Cach       PaymentType = "cach"
	CreditCard PaymentType = "credit_card"
)

const (
	Pending PaymentStatus = 1
	Partial PaymentStatus = 2
	Paid    PaymentStatus = 3
)

type CreateOrderModel struct {
	CompanyID   string  `json:"company_id" binding:"required"`
	ClientID    int     `json:"client_id"`
	Phone       string  `json:"phone" binding:"required"`
	Count       int     `json:"count"`
	Slug        string  `json:"slug"`
	Status      int8    `json:"status"`
	Description string  `json:"description"`
	ChatID      int64   `json:"chat_id"`
	Address     string  `json:"address" binding:"required"`
	IsNewClient bool    `json:"is_new_client"`
	Latitute    float64 `json:"latitute"`
	Longitude   float64 `json:"longitude"`
}

type OrderList struct {
	ID        int       `json:"id"`
	Slug      string    `json:"slug"`
	Phone     string    `json:"phone"`
	Address   *string   `json:"address"`
	Status    int16     `json:"status"`
	Square    float64   `json:"square"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderListResponse struct {
	Data  []OrderList `json:"data"`
	Count int         `json:"total"`
}

type OrdersListRequest struct {
	Limit         int32         `json:"limit" form:"limit"`
	Offset        int32         `json:"offset" form:"offset"`
	Status        int           `json:"status" form:"status"`
	PaymentStatus PaymentStatus `json:"payment_status" form:"payment_status"`
	Phone         string        `json:"phone" form:"phone"`
	ID            string        `json:"id" form:"id"`
	DateFrom      time.Time     `json:"date_from" form:"date_from"`
	DateTo        time.Time     `json:"date_to" form:"date_to"`
	CompanyID     string        `json:"company_id" form:"company_id" binding:"required"`
}

type OrderShowResponse struct {
	Order
	OrderItems       []OrderItem        `json:"order_items"`
	OrderTransaction []OrderTransaction `json:"transactions"`
}

type OrderTransaction struct {
	ReceiverFullname string    `json:"receiver_fullname"`
	PaymentType      string    `json:"payment_type"`
	Amount           float64   `json:"amount"`
	CreatedAt        time.Time `json:"created_at"`
}

type Order struct {
	ID                    int                   `json:"id"`
	CompanyID             string                `json:"company_id"`
	ClientID              int                   `json:"client_id"`
	PhoneNumber           string                `json:"phone_number"`
	AdditionalPhoneNumber string                `json:"additional_phone_number"`
	WorkNumber            string                `json:"work_number"`
	Count                 int                   `json:"count"`
	Slug                  string                `json:"slug"`
	Status                int8                  `json:"status"`
	Description           string                `json:"description"`
	CreatedAt             time.Time             `json:"created_at"`
	UpdatedAt             time.Time             `json:"updated_at"`
	Latitute              *float64              `json:"latitute"`
	Longitude             *float64              `json:"longitude"`
	Address               *string               `json:"address"`
	Square                float64               `json:"square"`
	Price                 float64               `json:"price"`
	StatusChangeHistory   []StatusChangeHistory `json:"status_change_history"`
	PaymentStatus         int16                 `json:"payment_status"`
	ServicePrice          float64               `json:"service_price"`
	DiscountPercentage    float64               `json:"discount_percentage"`
	DiscountPrice         float64               `json:"discounted_price"`
}

type OrderSendLocationRequest struct {
	OrderID int `form:"order_id" binding:"required"`
}

type UpdateOrderRequest struct {
	ID            int           `json:"id" binding:"required"`
	CompanyID     string        `json:"company_id"`
	Address       string        `json:"address"`
	Slug          string        `json:"slug"`
	Status        int8          `json:"status"`
	PaymentStatus PaymentStatus `json:"payment_status"`
	Phone         string        `json:"phone"`
	ChatID        int64         `json:"chat_id"`
	Description   string        `json:"description"`
	Count         int           `json:"count"`
	Latitute      float64       `json:"latitute"`
	Longitude     float64       `json:"longitude"`
}

type AddOrderPaymentRequest struct {
	CompanyID   string  `json:"company_id" binding:"required"`
	OrderID     int     `json:"order_id" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	PaymentType string  `json:"payment_type" binding:"required,oneof=cach credit_card"`
	Description string  `json:"description"`
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

type DeleteOrderRequest struct {
	ID        int    `json:"id" binding:"required"`
	CompanyID string `json:"company_id" binding:"required"`
}

type SetOrderPriceRequest struct {
	ID        int    `json:"order_id" binding:"required"`
	CompanyID string `json:"company_id" binding:"required"`
	// ServicePrice float64 `json:"service_price" binding:"required"`
	// DiscountPercentage float64 `json:"discount_percentage" binding:"required"`
	DiscountedPrice float64 `json:"discounted_price"`
}
