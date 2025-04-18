package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	if ipProxy := r.Header.Get("X-Forwarded-For"); ipProxy != "" {
		ip = ipProxy
	}
	fmt.Fprintln(w, ip)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
