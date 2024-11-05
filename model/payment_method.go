package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentMethod struct {
	ID           primitive.ObjectID `json:"payment_method_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`             // Nama metode pembayaran (e.g., Bank Transfer, E-Wallet)
	AdditionalFee float32           `json:"additional_fee,omitempty" bson:"additional_fee,omitempty"` // Biaya tambahan untuk metode pembayaran
	Discount     float32           `json:"discount,omitempty" bson:"discount,omitempty"`       // Diskon yang mungkin diterapkan
	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
