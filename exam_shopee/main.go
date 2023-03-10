package main

import (
	"fmt"
	"net/http"
)

func getClientIpAddr(req *http.Request) string {
	clientIp := req.Header.Get("X-FORWARDED-FOR")
	if clientIp != "" {
		return clientIp
	}
	return req.RemoteAddr
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	clientIp := getClientIpAddr(r)

	fmt.Fprintf(w, clientIp+"\n\n")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/health", HealthCheck)
	http.ListenAndServe(":8080", nil)
}