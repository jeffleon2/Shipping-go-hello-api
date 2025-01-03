package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jeffleon2/shipping-go-hello-api/handlers"
	"github.com/jeffleon2/shipping-go-hello-api/handlers/rest"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}

	server := &http.Server{
		Addr:         addr,
		Handler:      http.DefaultServeMux, // Puedes personalizar el handler aquí
		ReadTimeout:  10 * time.Second,     // Máximo tiempo para leer la solicitud
		WriteTimeout: 10 * time.Second,     // Máximo tiempo para escribir la respuesta
		IdleTimeout:  120 * time.Second,
	}

	http.HandleFunc("/hello", rest.TranslateHandler)
	http.HandleFunc("/health", handlers.HealthCheck)

	log.Printf("listening on %s\n", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
