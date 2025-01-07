package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello world")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "sorry", http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, "Hello ", string(d))
	})
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye")
		log.Fatal()
	})
	http.ListenAndServe("localhost:8080", nil)
}
