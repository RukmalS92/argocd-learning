package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Version is what we'll bump to watch ArgoCD roll out changes.
const Version = "1.0.0"

func main() {
	hostname, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from go-app!\nversion: %s\npod: %s\n", Version, hostname)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	log.Printf("go-app v%s listening on :8080", Version)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
