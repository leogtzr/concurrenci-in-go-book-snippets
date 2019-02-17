package main

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

func Open() *APIConnection {
	secondLimit := rate.NewLimiter(Per(2, time.Second), 1)
	minuteLimit := rate.NewLimiter(Per(10, time.Minute), 10)
	return &APIConnection{
		//rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
		rateLimiter: MultiLimiter(secondLimit, minuteLimit),
	}
}

// APIConnection ...
type APIConnection struct {
	rateLimiter RateLimiter
}

func (a *APIConnection) ReadFile(ctx context.Context) error {

	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}

	// Pretend we do work here ...
	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {

	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}

	// Pretend we do work here ...
	return nil
}
