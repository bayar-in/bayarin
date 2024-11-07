package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// merchant struct
type Merchant struct {
	ID          primitive.ObjectID `json:"merchant_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name_merchant,omitempty" bson:"name,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	Balance     float32            `json:"balance,omitempty" bson:"balance,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
