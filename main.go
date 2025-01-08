package main

import (
	"context"
	"log"
	"microservies/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hel := handlers.NewHello(l)
	godbye := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hel)
	sm.Handle("/goodbye", godbye)

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
