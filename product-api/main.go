package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kaushiknishant/go-microservices/product-api/handlers"
)

func main() {
	logs := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logs)
	goodByeHandler := handlers.NewGoodBye(logs)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodByeHandler)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logs.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logs.Println("Received terminate, graceful shutdown", sig)

	duration := time.Now().Add(30 * time.Second)
	timeOutContext, cancel := context.WithDeadline(context.Background(), duration)

	defer cancel()

	server.Shutdown(timeOutContext)
}
