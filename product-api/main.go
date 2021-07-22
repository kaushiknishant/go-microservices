package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kaushiknishant/go-microservices/product-api/handlers"
)

func main() {
	logs := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logs)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)

	http.ListenAndServe(":9090", serveMux)
}
