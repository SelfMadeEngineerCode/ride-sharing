package types

import "time"

// PaymentStatus represents the current status of a payment
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSuccess   PaymentStatus = "success"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

// Payment represents a payment transaction
type Payment struct {
	ID              string        `json:"id"`
	TripID          string        `json:"trip_id"`
	UserID          string        `json:"user_id"`
	Amount          int64         `json:"amount"`   // Amount in cents
	Currency        string        `json:"currency"` // e.g., "usd"
	Status          PaymentStatus `json:"status"`
	StripeSessionID string        `json:"stripe_session_id"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

// PaymentIntent represents the intent to collect a payment
type PaymentIntent struct {
	ID              string    `json:"id"`
	TripID          string    `json:"trip_id"`
	UserID          string    `json:"user_id"`
	DriverID        string    `json:"driver_id"`
	Amount          int64     `json:"amount"`
	Currency        string    `json:"currency"`
	StripeSessionID string    `json:"stripe_session_id"`
	CreatedAt       time.Time `json:"created_at"`
}

// PaymentConfig holds the configuration for the payment service
type PaymentConfig struct {
	StripeSecretKey      string `json:"stripeSecretKey"`
	StripePublishableKey string `json:"stripePublishableKey"`
	StripeWebhookSecret  string `json:"stripeWebhookSecret"`
	Currency             string `json:"currency"`
	SuccessURL           string `json:"successURL"`
	CancelURL            string `json:"cancelURL"`
}
