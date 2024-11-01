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

// CreatePembayaran untuk menambahkan pembayaran baru
func AddpembayaranToorder(w http.ResponseWriter, r *http.Request) {
	// Mengambil slug dari query parameter
	slug := r.URL.Query().Get("slug")

	// Dekode body permintaan untuk mendapatkan detail pembayaran baru
	var newpembayaran model.Payment
	err := json.NewDecoder(r.Body).Decode(&newpembayaran)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set id pembayaran baru dan waktu pembuatan
	newpembayaran.ID = primitive.NewObjectID()       // Set ObjectID baru untuk pembayaran
	newpembayaran.CreatedAt = time.Now()             // Set waktu saat ini untuk createdAt
	newpembayaran.UpdatedAt = time.Now()             // Set waktu saat ini untuk updatedAt

	// Membuat filter untuk mencocokkan order berdasarkan slug
	filter := bson.M{"slug": slug}

	update := bson.M{
		"$push": bson.M{
			"pembayaran": newpembayaran,
		},
	}

	collection := config.Mongoconn.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Failed to add pembayaran", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	// Kirim respons berhasil
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pembayaran added successfully"}`))
}

// GetPembayaran untuk mendapatkan detail pembayaran
func Getpembayaran(w http.ResponseWriter, r *http.Request) {
	// Mengambil slug dari query parameter
	slug := r.URL.Query().Get("slug")

	// Membuat filter untuk mencocokkan order berdasarkan slug
	filter := bson.M{"slug": slug}

	collection := config.Mongoconn.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order model.Order
	err := collection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	// Kirim respons dengan detail order
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// UpdatePembayaran untuk memperbarui detail pembayaran
func Updatepembayaran(w http.ResponseWriter, r *http.Request) {
	// Mengambil slug dari query parameter
	slug := r.URL.Query().Get("slug")

	// Dekode body permintaan untuk mendapatkan detail pembayaran baru
	var updatedpembayaran model.Payment
	err := json.NewDecoder(r.Body).Decode(&updatedpembayaran)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set waktu pembaharuan
	updatedpembayaran.UpdatedAt = time.Now()

	// Membuat filter untuk mencocokkan order berdasarkan slug dan id pembayaran
	filter := bson.M{
		"slug": slug,
		"pembayaran._id": updatedpembayaran.ID,
	}

	update := bson.M{
		"$set": bson.M{
			"pembayaran.$": updatedpembayaran,
		},
	}

	collection := config.Mongoconn.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Failed to update pembayaran", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "order or pembayaran not found", http.StatusNotFound)
		return
	}

	// Kirim respons berhasil
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pembayaran updated successfully"}`))
}

