package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Notifikasi struct
type Notifikasi struct {
	ID            primitive.ObjectID `json:"notifikasi_id,omitempty" bson:"_id,omitempty"`
	UserID        primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	TransactionID primitive.ObjectID `json:"transaction_id,omitempty" bson:"transaction_id,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	Message       string             `json:"message,omitempty" bson:"message,omitempty"`
	CreatedAt     time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
