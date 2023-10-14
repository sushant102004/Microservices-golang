package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	microservices "github.com/sushant102004/microservices/proto"
	"github.com/sushant102004/microservices/types"
	"google.golang.org/grpc"
)

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, crypto string) (*types.PriceResonse, error) {
	c.endpoint = fmt.Sprintf("%s?crypto=%s", c.endpoint, crypto)

	req, err := http.NewRequest("get", c.endpoint, nil)
	if err != nil {
		return nil, nil
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		httpErr := make(map[string]any)
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error %s", httpErr["error"])
	}

	price := new(types.PriceResonse)

	if err := json.NewDecoder(resp.Body).Decode(&price); err != nil {
		return nil, err
	}
	return price, nil
}

func NewGRPCClient(targetAddr string) (microservices.PriceFetcherClient, error) {
	conn, err := grpc.Dial(targetAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return microservices.NewPriceFetcherClient(conn), nil
}
