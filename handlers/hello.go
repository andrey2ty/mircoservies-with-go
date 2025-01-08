package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	lg *log.Logger
}

func NewHello(lg *log.Logger) *Hello {
	return &Hello{lg}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.lg.Println("hello world")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "sorry", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Hello ", string(d))
}
