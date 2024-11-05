package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gocroot/config"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddMerchant untuk menambahkan merchant baru
func AddMerchant(w http.ResponseWriter, r *http.Request) {
	var merchant model.Merchant
	if err := json.NewDecoder(r.Body).Decode(&merchant); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	merchant.ID = primitive.NewObjectID()
	merchant.CreatedAt = time.Now()
	merchant.UpdatedAt = time.Now()

	collection := config.Mongoconn.Collection("merchants")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, merchant)
	if err != nil {
		http.Error(w, "Failed to add merchant", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(merchant)
}

// GetMerchants untuk mendapatkan semua merchant
func GetMerchants(w http.ResponseWriter, r *http.Request) {
	collection := config.Mongoconn.Collection("merchants")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch merchants", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var merchants []model.Merchant
	if err = cursor.All(ctx, &merchants); err != nil {
		http.Error(w, "Failed to decode merchants", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(merchants)
}

// GetMerchantByID untuk mendapatkan merchant berdasarkan ID
func GetMerchantByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	collection := config.Mongoconn.Collection("merchants")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var merchant model.Merchant
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&merchant)
	if err != nil {
		http.Error(w, "Failed to fetch merchant", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(merchant)
}

// UpdateMerchant untuk memperbarui merchant yang ada
func UpdateMerchant(w http.ResponseWriter, r *http.Request) {
	var merchant model.Merchant
	if err := json.NewDecoder(r.Body).Decode(&merchant); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	merchant.UpdatedAt = time.Now()

	collection := config.Mongoconn.Collection("merchants")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.ReplaceOne(ctx, bson.M{"_id": merchant.ID}, merchant)
	if err != nil {
		http.Error(w, "Failed to update merchant", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(merchant)
}

// DeleteMerchant untuk menghapus merchant berdasarkan ID
func DeleteMerchant(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	collection := config.Mongoconn.Collection("merchants")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		http.Error(w, "Failed to delete merchant", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}