/*
Package retry provides a simple retry mechanism with exponential backoff.
It is as abstract as possible to allow for different retry strategies.
*/
package retry

import (
	"context"
	"log"
	"time"
)

type Config struct {
	MaxRetries  int
	InitialWait time.Duration
	MaxWait     time.Duration
}

// DefaultConfig returns a Config with sensible default values
func DefaultConfig() Config {
	return Config{
		MaxRetries:  3,
		InitialWait: 1 * time.Second,
		MaxWait:     10 * time.Second,
	}
}

// WithBackoff executes the given operation with exponential backoff retry logic
func WithBackoff(ctx context.Context, cfg Config, operation func() error) error {
	var err error
	wait := cfg.InitialWait

	for attempt := 0; attempt <= cfg.MaxRetries; attempt++ {
		if attempt > 0 {
			log.Printf("Retry attempt %d/%d after %v", attempt, cfg.MaxRetries, wait)

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(wait):
			}

			// Exponential backoff with max wait cap
			wait *= 2
			if wait > cfg.MaxWait {
				wait = cfg.MaxWait
			}
		}

		if err = operation(); err == nil {
			return nil
		}

		log.Printf("Operation failed (attempt %d/%d): %v", attempt+1, cfg.MaxRetries, err)
	}

	return err
}
