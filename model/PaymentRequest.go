package model

type PaymentRequest struct {
	Jumlah        int    `json:"jumlah"`
	PaymentMethod string `json:"payment_method"`
	OrderID       string `json:"order_id"`
}

type PaymentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
