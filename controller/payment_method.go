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

// AddPaymentMethod untuk menambahkan metode pembayaran baru
func AddPaymentMethod(w http.ResponseWriter, r *http.Request) {
	var paymentMethod model.PaymentMethod
	if err := json.NewDecoder(r.Body).Decode(&paymentMethod); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	paymentMethod.ID = primitive.NewObjectID()
	paymentMethod.CreatedAt = time.Now()
	paymentMethod.UpdatedAt = time.Now()

	collection := config.Mongoconn.Collection("payment_methods")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, paymentMethod)
	if err != nil {
		http.Error(w, "Failed to add payment method", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(paymentMethod)
}

// GetPaymentMethods untuk mendapatkan semua metode pembayaran
func GetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	collection := config.Mongoconn.Collection("payment_methods")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch payment methods", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var paymentMethods []model.PaymentMethod
	if err = cursor.All(ctx, &paymentMethods); err != nil {
		http.Error(w, "Failed to decode payment methods", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(paymentMethods)
}

// GetPaymentMethodByID untuk mendapatkan metode pembayaran berdasarkan ID
func GetPaymentMethodByID(w http.ResponseWriter, r *http.Request) {
	collection := config.Mongoconn.Collection("payment_methods")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := r.URL.Query().Get("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var paymentMethod model.PaymentMethod
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&paymentMethod)
	if err != nil {
		http.Error(w, "Failed to fetch payment method", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(paymentMethod)
}

// UpdatePaymentMethod untuk mengubah data metode pembayaran
func UpdatePaymentMethod(w http.ResponseWriter, r *http.Request) {
	collection := config.Mongoconn.Collection("payment_methods")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := r.URL.Query().Get("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var paymentMethod model.PaymentMethod
	if err := json.NewDecoder(r.Body).Decode(&paymentMethod); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	paymentMethod.UpdatedAt = time.Now()

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": paymentMethod})
	if err != nil {
		http.Error(w, "Failed to update payment method", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(paymentMethod)
}

// DeletePaymentMethod untuk menghapus metode pembayaran
func DeletePaymentMethod(w http.ResponseWriter, r *http.Request) {
	collection := config.Mongoconn.Collection("payment_methods")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := r.URL.Query().Get("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		http.Error(w, "Failed to delete payment method", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
