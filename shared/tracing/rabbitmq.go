package tracing

import (
	"context"
	"encoding/json"

	"ride-sharing/shared/contracts"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// amqpHeadersCarrier implements the TextMapCarrier interface for AMQP headers
type amqpHeadersCarrier amqp.Table

func (c amqpHeadersCarrier) Get(key string) string {
	if v, ok := c[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func (c amqpHeadersCarrier) Set(key string, value string) {
	c[key] = value
}

func (c amqpHeadersCarrier) Keys() []string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}

// TracedPublisher wraps the RabbitMQ publish function with tracing
func TracedPublisher(ctx context.Context, exchange, routingKey string, msg amqp.Publishing, publish func(context.Context, string, string, amqp.Publishing) error) error {
	tracer := otel.GetTracerProvider().Tracer("rabbitmq")

	ctx, span := tracer.Start(ctx, "rabbitmq.publish",
		trace.WithAttributes(
			attribute.String("messaging.destination", exchange),
			attribute.String("messaging.routing_key", routingKey),
		),
	)
	defer span.End()

	// Try to extract and add message details to span (map[string]any if you don't know the type)
	var msgBody contracts.AmqpMessage
	if err := json.Unmarshal(msg.Body, &msgBody); err == nil {
		if msgBody.OwnerID != "" {
			span.SetAttributes(attribute.String("messaging.owner_id", msgBody.OwnerID))
		}
	}

	// Inject trace context into message headers
	if msg.Headers == nil {
		msg.Headers = make(amqp.Table)
	}
	carrier := amqpHeadersCarrier(msg.Headers)
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	msg.Headers = amqp.Table(carrier)

	if err := publish(ctx, exchange, routingKey, msg); err != nil {
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	return nil
}

// TracedConsumer wraps the RabbitMQ message handler with tracing
func TracedConsumer(delivery amqp.Delivery, handler func(context.Context, amqp.Delivery) error) error {
	// Extract trace context from message headers
	carrier := amqpHeadersCarrier(delivery.Headers)
	ctx := otel.GetTextMapPropagator().Extract(context.Background(), carrier)

	tracer := otel.GetTracerProvider().Tracer("rabbitmq")

	ctx, span := tracer.Start(ctx, "rabbitmq.consume",
		trace.WithAttributes(
			attribute.String("messaging.destination", delivery.Exchange),
			attribute.String("messaging.routing_key", delivery.RoutingKey),
		),
	)
	defer span.End()

	// Try to extract and add message details to span (map[string]any if you don't know the type)
	var msgBody contracts.AmqpMessage
	if err := json.Unmarshal(delivery.Body, &msgBody); err == nil {
		if msgBody.OwnerID != "" {
			span.SetAttributes(attribute.String("messaging.owner_id", msgBody.OwnerID))
		}
	}

	if err := handler(ctx, delivery); err != nil {
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	return nil
}
