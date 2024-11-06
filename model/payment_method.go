package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentMethod struct {
	ID           primitive.ObjectID `json:"payment_method_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`             // Nama metode pembayaran (e.g., Bank Transfer, E-Wallet)
	Description  string             `json:"description,omitempty" bson:"description,omitempty"` // Deskripsi metode pembayaran
	Type 	   string             `json:"type,omitempty" bson:"type,omitempty"`             // Tipe metode pembayaran (e.g., Bank Transfer, E-Wallet)
	Provider	 string             `json:"provider,omitempty" bson:"provider,omitempty"`     // Provider metode pembayaran (e.g., BCA, OVO)
	Fee 		float32            `json:"fee,omitempty" bson:"fee,omitempty"`               // Biaya transaksi
	Currency	 string             `json:"currency,omitempty" bson:"currency,omitempty"`     // Mata uang
	Status	   string             `json:"status,omitempty" bson:"status,omitempty"`         // Status metode pembayaran (e.g., Active, Inactive)
	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
