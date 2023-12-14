package models

import "time"

type Order struct {
	OrderUid          string    `json:"order_uid" validate:"required,min=19,max=19"`
	TrackNumber       string    `json:"track_number" validate:"required,min=14,max=14"`
	Entry             string    `json:"entry" validate:"required"`
	Locale            string    `json:"locale" validate:"required,oneof=ru en"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" validate:"required"`
	ShardKey          string    `json:"shardkey" validate:"required"`
	SmId              int       `json:"sm_id" validate:"required"`
	DateCreated       time.Time `json:"date_created" format:"2006-01-02T06:22:19Z" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []Item    `json:"items" validate:"required"`
}
