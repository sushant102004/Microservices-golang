package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sushant102004/microservices/client"
)

func main() {
	svc := NewLoggingService(&priceFetcher{})
	server := NewJSONServer(svc, ":3000")
	server.Run()

	client := client.New("http://localhost:3000/")
	price, err := client.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal("Error - ", err.Error())
	}

	fmt.Println(price)
}
