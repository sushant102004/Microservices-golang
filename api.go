/*
	Create and start HTTP listener.
*/

package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/sushant102004/microservices/types"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type JSONServer struct {
	listenAddr string
	svc        PriceFetcher
}

func NewJSONServer(svc PriceFetcher, listenAddr string) *JSONServer {
	return &JSONServer{
		svc:        svc,
		listenAddr: listenAddr,
	}
}

func (s *JSONServer) Run() {
	// We can't use s.handleGetPrice directly because it do not obey the conventions of standard http handler function.
	// So we will write a function that will accept another function and convert it into standard http function.

	http.HandleFunc("/", newHTTPHandlerFunc(s.handleGetPrice))
	http.ListenAndServe(s.listenAddr, nil)
}

func newHTTPHandlerFunc(apiFunc APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "request-id", rand.Intn(10000000000))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(ctx, w, r); err != nil {
			writeJson(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
		}
	}
}

func (s *JSONServer) handleGetPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	crypto := r.URL.Query().Get("crypto")

	price, err := s.svc.FetchPrice(ctx, crypto)
	if err != nil {
		return err
	}

	res := &types.PriceResonse{
		Crypto: crypto,
		Price:  price,
	}

	return writeJson(w, http.StatusOK, res)
}

func writeJson(w http.ResponseWriter, s int, body any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(body)
}
