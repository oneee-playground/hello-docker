package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello Docker!"))
	}))

	http.Handle("/vary", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"foo":"bar"}`))
		} else {
			w.WriteHeader(http.StatusTeapot)
		}
	}))

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(":4000", nil); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	log.Println("Server Started!")
	wg.Wait()
}
