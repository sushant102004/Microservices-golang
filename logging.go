/*
	PoF - Implement decorator pattern and wrap core FetchPrice logic over logging service to print logs on every request.
*/

package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggingService struct {
	next PriceFetcher
}

// We are able to return LoggingService because that struct implement FetchPrice function which is also implemented by PriceFetcher
// which makes both of them compatible.
func NewLoggingService(next PriceFetcher) PriceFetcher {
	return &LoggingService{
		next: next,
	}
}

func (l *LoggingService) FetchPrice(ctx context.Context, crypto string) (price float64, err error) {
	defer func(sTime time.Time) {
		logrus.WithFields(logrus.Fields{
			"request-id": ctx.Value("request-id"),
			"took":       time.Since(sTime),
			"error":      err,
			"price":      price,
		}).Info("Fetch Price")
	}(time.Now())

	return l.next.FetchPrice(ctx, crypto)
}
