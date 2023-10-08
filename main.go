package main

func main() {
	svc := NewLoggingService(&priceFetcher{})
	server := NewJSONServer(svc, ":3000")
	server.Run()
}
