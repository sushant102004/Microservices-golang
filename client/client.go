package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sushant102004/microservices/types"
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
