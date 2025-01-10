package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	handlers2 "microservies/product-api/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	pr := handlers2.NewProducts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", pr.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", pr.UpdateProduct)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	go func() {
		s.ListenAndServe()
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
