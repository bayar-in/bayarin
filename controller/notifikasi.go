package controller

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "github.com/gocroot/config"
    "github.com/gocroot/helper/services"
    "github.com/gocroot/model"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// AddNotifikasi untuk menambahkan notifikasi baru
func AddNotifikasi(w http.ResponseWriter, r *http.Request) {
    var notifikasi model.Notifikasi
    if err := json.NewDecoder(r.Body).Decode(&notifikasi); err != nil {
        http.Error(w, "Invalid request data: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Validasi tambahan untuk data notifikasi
    if notifikasi.Email == "" || notifikasi.Message == "" {
        http.Error(w, "Email and Message fields are required", http.StatusBadRequest)
        return
    }

    // Inisialisasi nilai ID dan waktu
    notifikasi.ID = primitive.NewObjectID()
    notifikasi.CreatedAt = time.Now()
    notifikasi.UpdatedAt = time.Now()

    collection := config.Mongoconn.Collection("notifikasi")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Memasukkan data ke MongoDB
    _, err := collection.InsertOne(ctx, notifikasi)
    if err != nil {
        http.Error(w, "Failed to add notifikasi: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Mengirim notifikasi email
    err = services.SendPaymentNotification(notifikasi.Email, "Notifikasi Baru", "Anda memiliki notifikasi baru.")
    if err != nil {
        http.Error(w, "Failed to send email notification: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(notifikasi)
}
