package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

// getClientIP extracts the real client IP and status of X-Forwarded-For
func getClientIP(r *http.Request) (tcpIP string, realIP string, headerSet bool) {
	tcpIP, _, _ = net.SplitHostPort(r.RemoteAddr)

	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		return tcpIP, forwardedFor, true
	}
	return tcpIP, tcpIP, false
}

// commonInfo writes base IP info to the response and logs it
func commonInfo(w http.ResponseWriter, r *http.Request) (tcpIP string, realIP string, headerSet bool) {
	tcpIP, realIP, headerSet = getClientIP(r)

	log.Printf("🔹 TCP Source IP: %s\n", tcpIP)
	log.Printf("🔸 Real Client IP: %s\n", realIP)

	fmt.Fprintf(w, "TCP Source IP     : %s\n", tcpIP)
	fmt.Fprintf(w, "Client Real IP    : %s\n", realIP)
	if headerSet {
		fmt.Fprintln(w, "X-Forwarded-For   : set")
	} else {
		fmt.Fprintln(w, "X-Forwarded-For   : not set")
	}

	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	commonInfo(w, r)
}

func handlerDetail(w http.ResponseWriter, r *http.Request) {
	commonInfo(w, r)

	fmt.Fprintln(w, "\n=== Request Headers ===")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
}

func main() {
	port := getEnv("PORT", "8080")
	addr := ":" + port
	log.Printf("✅ Server starting on %s\n", addr)

	http.HandleFunc("/", handler)
	http.HandleFunc("/detail", handlerDetail)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
