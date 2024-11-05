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

// controller/pembayaran.go

func AddPembayaranToOrder(w http.ResponseWriter, r *http.Request) {
	// Ambil ID order dari query parameter
	orderIDHex := r.URL.Query().Get("order_id")
	orderID, err := primitive.ObjectIDFromHex(orderIDHex)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	// Decode body permintaan untuk mendapatkan detail pembayaran baru
	var newPembayaran model.Payment
	err = json.NewDecoder(r.Body).Decode(&newPembayaran)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set OrderID dan waktu pembuatan
	newPembayaran.OrderID = orderID
	newPembayaran.ID = primitive.NewObjectID()
	newPembayaran.CreatedAt = time.Now()
	newPembayaran.UpdatedAt = time.Now()

	collection := config.Mongoconn.Collection("pembayaran")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, newPembayaran)
	if err != nil {
		http.Error(w, "Failed to add payment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Payment added successfully"}`))
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

// controller/pembayaran.go

func GetPembayaranByOrderID(w http.ResponseWriter, r *http.Request) {
	orderIDHex := r.URL.Query().Get("order_id")
	orderID, err := primitive.ObjectIDFromHex(orderIDHex)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	filter := bson.M{"order_id": orderID}

	collection := config.Mongoconn.Collection("pembayaran")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, "Failed to fetch payments", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var payments []model.Payment
	if err = cursor.All(ctx, &payments); err != nil {
		http.Error(w, "Failed to decode payments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

// controller/pembayaran.go
func GetPembayaranWithMethod(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID pembayaran dari query parameter
	paymentIDHex := r.URL.Query().Get("payment_id")
	paymentID, err := primitive.ObjectIDFromHex(paymentIDHex)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	collection := config.Mongoconn.Collection("pembayaran")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var payment model.Payment
	err = collection.FindOne(ctx, bson.M{"_id": paymentID}).Decode(&payment)
	if err != nil {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}

	// Ambil metode pembayaran berdasarkan PaymentMethodID
	methodCollection := config.Mongoconn.Collection("payment_methods")
	var paymentMethod model.PaymentMethod
	err = methodCollection.FindOne(ctx, bson.M{"_id": payment.PaymentMethodID}).Decode(&paymentMethod)
	if err != nil {
		http.Error(w, "Payment method not found", http.StatusNotFound)
		return
	}

	// Kirim respons dengan detail pembayaran dan metode pembayaran
	response := struct {
		Payment       model.Payment       `json:"payment"`
		PaymentMethod model.PaymentMethod  `json:"payment_method"`
	}{
		Payment:       payment,
		PaymentMethod: paymentMethod,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
