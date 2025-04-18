package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	if ipProxy := r.Header.Get("X-Forwarded-For"); ipProxy != "" {
		ip = ipProxy
	}

	// Log the response
	log.Printf("Client IP: %s\n", ip)

	fmt.Fprintln(w, ip)
}

func main() {

	port := getEnv("PORT", "8080")
	addr := ":" + port
	log.Printf("✅ Server starting on %s\n", addr)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
