package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sushant102004/microservices/client"
	microservices "github.com/sushant102004/microservices/proto"
)

func main() {
	var (
		svc = NewLoggingService(&priceFetcher{})
		ctx = context.Background()
	)

	go makeGRPCServer(":4000", svc)

	go func() {
		time.Sleep(time.Second * 3)
		client := client.New("http://localhost:3000/")
		price, err := client.FetchPrice(context.Background(), "ETH")
		if err != nil {
			log.Fatal("Error - ", err.Error())
		}
		fmt.Println("**JSON Response**")
		fmt.Println(price)
	}()

	go func() {
		client, err := client.NewGRPCClient(":4000")
		if err != nil {
			log.Fatal(err)
		}

		resp, err := client.FetchPrice(ctx, &microservices.PriceRequest{
			Crypto: "ETH",
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("**GRPC Response**")
		fmt.Println(resp)
	}()

	server := NewJSONServer(svc, ":3000")
	server.Run()
}
