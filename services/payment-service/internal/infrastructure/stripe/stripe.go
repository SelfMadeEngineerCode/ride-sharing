package stripe

import (
	"context"
	"fmt"
	"ride-sharing/services/payment-service/internal/domain"
	"ride-sharing/services/payment-service/pkg/types"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
)

type stripeClient struct {
	config *types.PaymentConfig
}

func NewStripeClient(config *types.PaymentConfig) domain.PaymentProcessor {
	stripe.Key = config.StripeSecretKey

	return &stripeClient{
		config: config,
	}
}

func (s *stripeClient) CreatePaymentSession(ctx context.Context, amount int64, currency string, metadata map[string]string) (string, error) {
	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String(s.config.SuccessURL),
		CancelURL:  stripe.String(s.config.CancelURL),
		Metadata: metadata,
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(currency),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Ride Payment"),
					},
					UnitAmount: stripe.Int64(amount),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
	}

	result, err := session.New(params)
	if err != nil {
		return "", fmt.Errorf("failed to create a payment session on stripe: %w", err)
	}

	return result.ID, nil
}
