package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/kaushiknishant/go-microservices/product-api/handlers"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	env.Parse()

	logs := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handlers
	productsHandler := handlers.NewProducts(logs)

	// create a new serve mux and register the handlers
	serveMux := mux.NewRouter()
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	//serveMux.Handle("/products", productsHandler)
	getRouter.HandleFunc("/", productsHandler.GetProducts)
	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productsHandler.UpdateProduct)

	postRouter := serveMux.Methods(http.MethodPut).Subrouter()
	postRouter.HandleFunc("/{id:[0-9]+}", productsHandler.AddProduct)
	// create a new server
	server := &http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      serveMux,          // set the default handler
		ErrorLog:     logs,              // set the logger for the server
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
	}

	// start the server
	go func() {
		logs.Println("Starting server on port 9090")

		err := server.ListenAndServe()
		if err != nil {
			logs.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received.
	sig := <-sigChan
	logs.Println("Received terminate, graceful shutdown", sig)

	// Absolute time needed for WithDeadline
	duration := time.Now().Add(30 * time.Second)
	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	timeOutContext, cancel := context.WithDeadline(context.Background(), duration)

	// Even though timeOutContext will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	// shutdown the server
	server.Shutdown(timeOutContext)
}
