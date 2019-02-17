package main

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

func Open() *APIConnection {
	//secondLimit := rate.NewLimiter(Per(2, time.Second), 1)
	//minuteLimit := rate.NewLimiter(Per(10, time.Minute), 10)
	return &APIConnection{
		//rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
		//rateLimiter: MultiLimiter(secondLimit, minuteLimit),
		apiLimit: MultiLimiter(
			rate.NewLimiter(Per(2, time.Second), 2),
			rate.NewLimiter(Per(10, time.Minute), 10),
		),
		diskLimit: MultiLimiter(
			rate.NewLimiter(rate.Limit(1), 1),
		),
		networkLimit: MultiLimiter(
			rate.NewLimiter(Per(3, time.Second), 3),
		),
	}
}

// APIConnection ...
type APIConnection struct {
	//rateLimiter RateLimiter
	networkLimit,
	diskLimit,
	apiLimit RateLimiter
}

func (a *APIConnection) ReadFile(ctx context.Context) error {

	// if err := a.rateLimiter.Wait(ctx); err != nil {
	// 	return err
	// }
	err := MultiLimiter(a.apiLimit, a.diskLimit).Wait(ctx)
	if err != nil {
		return err
	}

	// Pretend we do work here ...
	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {

	// if err := a.rateLimiter.Wait(ctx); err != nil {
	// 	return err
	// }
	err := MultiLimiter(a.apiLimit, a.networkLimit).Wait(ctx)
	if err != nil {
		return err
	}

	// Pretend we do work here ...
	return nil
}
