package main

import (
    "time"
)


type ModelJson struct {
	OrderUID    string `json:"order_uid,omitempty"`
	TrackNumber string `json:"track_number,omitempty"`
	Entry       string `json:"entry,omitempty"`
	Delivery    DeliveryType `json:"delivery,omitempty"`
	Payment PaymentType `json:"payment,omitempty"`
	Items ItemsType `json:"items,omitempty"`
	Locale            string    `json:"locale,omitempty"`
	InternalSignature string    `json:"internal_signature,omitempty"`
	CustomerID        string    `json:"customer_id,omitempty"`
	DeliveryService   string    `json:"delivery_service,omitempty"`
	Shardkey          string    `json:"shardkey,omitempty"`
	SmID              int       `json:"sm_id,omitempty"`
	DateCreated       time.Time `json:"date_created,omitempty"`
	OofShard          string    `json:"oof_shard,omitempty"`
}
type parseSrtuct struct {
    OrderUID string `json:"order_uid,omitempty"`
}
type DeliveryType struct {
    Name    string `json:"name,omitempty"`
    Phone   string `json:"phone,omitempty"`
    Zip     string `json:"zip,omitempty"`
    City    string `json:"city,omitempty"`
    Address string `json:"address,omitempty"`
    Region  string `json:"region,omitempty"`
    Email   string `json:"email,omitempty"`
}
type PaymentType struct {
    Transaction  string `json:"transaction,omitempty"`
    RequestID    string `json:"request_id,omitempty"`
    Currency     string `json:"currency,omitempty"`
    Provider     string `json:"provider,omitempty"`
    Amount       int    `json:"amount,omitempty"`
    PaymentDt    int    `json:"payment_dt,omitempty"`
    Bank         string `json:"bank,omitempty"`
    DeliveryCost int    `json:"delivery_cost,omitempty"`
    GoodsTotal   int    `json:"goods_total,omitempty"`
    CustomFee    int    `json:"custom_fee,omitempty"`
}
type ItemsType []Item
type Item struct {
    ChrtID      int    `json:"chrt_id,omitempty"`
    TrackNumber string `json:"track_number,omitempty"`
    Price       int    `json:"price,omitempty"`
    Rid         string `json:"rid,omitempty"`
    Name        string `json:"name,omitempty"`
    Sale        int    `json:"sale,omitempty"`
    Size        string `json:"size,omitempty"`
    TotalPrice  int    `json:"total_price,omitempty"`
    NmID        int    `json:"nm_id,omitempty"`
    Brand       string `json:"brand,omitempty"`
    Status      int    `json:"status,omitempty"`
}
