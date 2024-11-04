package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gocroot/config"
	"github.com/gocroot/helper/payment"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandlePayment(w http.ResponseWriter, r *http.Request) {
	var paymentRequest model.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&paymentRequest); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Lakukan proses pembayaran dengan fungsi yang sesuai di helper/payment
	result, err := payment.ProcessPayment(paymentRequest)
	if err != nil {
		http.Error(w, "Payment processing failed", http.StatusInternalServerError)
		return
	}

	// Berikan respon sukses jika berhasil
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// AddpembayaranToorder untuk menambahkan pembayaran baru ke dalam order
func AddpembayaranToorder(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID order dari query parameter
	orderID := r.URL.Query().Get("id")
	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	// Dekode body permintaan untuk mendapatkan detail pembayaran baru
	var newpembayaran model.Payment
	if err := json.NewDecoder(r.Body).Decode(&newpembayaran); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set id pembayaran baru dan waktu pembuatan
	newpembayaran.ID = primitive.NewObjectID() // Set ObjectID baru untuk pembayaran
	newpembayaran.CreatedAt = time.Now()       // Set waktu saat ini untuk createdAt
	newpembayaran.UpdatedAt = time.Now()       // Set waktu saat ini untuk updatedAt

	// Membuat filter untuk mencocokkan order berdasarkan ObjectID
	filter := bson.M{"_id": objID}
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

// Getpembayaran untuk mendapatkan detail pembayaran dari sebuah order
func Getpembayaran(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID order dari query parameter
	orderID := r.URL.Query().Get("id")
	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	// Membuat filter untuk mencocokkan order berdasarkan ObjectID
	filter := bson.M{"_id": objID}

	collection := config.Mongoconn.Collection("order")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order model.Order
	err = collection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	// Kirim respons dengan detail order
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// Updatepembayaran untuk memperbarui detail pembayaran dalam sebuah order
func Updatepembayaran(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID order dari query parameter
	orderID := r.URL.Query().Get("id")
	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	// Dekode body permintaan untuk mendapatkan detail pembayaran yang diperbarui
	var updatedpembayaran model.Payment
	if err := json.NewDecoder(r.Body).Decode(&updatedpembayaran); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set waktu pembaharuan
	updatedpembayaran.UpdatedAt = time.Now()

	// Membuat filter untuk mencocokkan order berdasarkan ObjectID dan id pembayaran
	filter := bson.M{
		"_id":            objID,
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
