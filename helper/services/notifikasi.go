package services

import (
	"fmt"
	"net/smtp"
)

// SendPaymentNotifikasi untuk mengirim notifikasi pembayaran
func SendPaymentNotification(email string, subject string, body string) error {
    smtpServer := "smtp.example.com"
    auth := smtp.PlainAuth("", "your-email@example.com", "your-password", smtpServer)

    msg := []byte("To: " + email + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "\r\n" +
        body + "\r\n")

    err := smtp.SendMail(smtpServer+":587", auth, "your-email@example.com", []string{email}, msg)
    if err != nil {
        return err
    }
    fmt.Println("Email sent to:", email)
    return nil
}