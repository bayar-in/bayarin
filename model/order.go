package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// transaction struct
type Transaction struct {
	ID              primitive.ObjectID `json:"transaction_id,omitempty" bson:"_id,omitempty"`
	UserID          primitive.ObjectID `json:"_id,omitempty" bson:"user_id,omitempty"`
	MerchantID      primitive.ObjectID `json:"merchant_id,omitempty" bson:"merchant_id,omitempty"`
	PaymentMethodID primitive.ObjectID `json:"payment_method_id,omitempty" bson:"payment_method_id,omitempty"`
	Amount          float32            `json:"amount,omitempty" bson:"amount,omitempty"`
	Currency        string             `json:"currency,omitempty" bson:"currency,omitempty"`
	Status          string             `json:"status,omitempty" bson:"status,omitempty"`
	Description     string             `json:"description,omitempty" bson:"description,omitempty"`
	CreatedAt       time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt       time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
