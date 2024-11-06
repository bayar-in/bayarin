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

// HandleTransaction untuk menangani transaksi
func HandleTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction model.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	transaction.ID = primitive.NewObjectID()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	collection := config.Mongoconn.Collection("transaction")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, transaction)
	if err != nil {
		http.Error(w, "Failed to add transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

// GetTransaction
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	collection := config.Mongoconn.Collection("transaction")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to get transaction", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var transactions []model.Transaction
	if err := cursor.All(ctx, &transactions); err != nil {
		http.Error(w, "Failed to get transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

// GetTransactionByID
func GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	collection := config.Mongoconn.Collection("transaction")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	var transaction model.Transaction
	if err := collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&transaction); err != nil {
		http.Error(w, "Failed to get transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

// UpdateTransaction
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction model.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	collection := config.Mongoconn.Collection("transaction")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	transaction.UpdatedAt = time.Now()
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": transaction})
	if err != nil {
		http.Error(w, "Failed to update transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

// DeleteTransaction
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	collection := config.Mongoconn.Collection("transaction")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, "Failed to delete transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}