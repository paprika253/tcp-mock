package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type Handler struct {
	mu *sync.Mutex
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	fmt.Println(time.Now().Format(time.RFC3339), "____________")
	fmt.Println(r.Method, r.Proto, r.URL.String())

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	fmt.Println(string(body))
	fmt.Println("____________________________________")
}

func main() {
	h := &Handler{
		&sync.Mutex{},
	}

	log.Fatal(http.ListenAndServe(":80", h))
}
