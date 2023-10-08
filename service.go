/*
	PoF - This will provide core logic of microservice.
*/

package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

var priceData = map[string]float64{
	"BTC": 32000,
	"ETH": 200,
}

func (s *priceFetcher) FetchPrice(ctx context.Context, crypto string) (float64, error) {
	return CryptoPriceFetcher(ctx, crypto)
}

func CryptoPriceFetcher(ctx context.Context, crypto string) (float64, error) {
	price, ok := priceData[crypto]
	if !ok {
		return price, fmt.Errorf("price not found for %s", crypto)
	}

	return price, nil
}
