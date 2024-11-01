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

