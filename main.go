package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

type echoResponse struct {
	Method  string            `json:"method"`
	Path    string            `json:"path"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body,omitempty"`
}

func clientIP(r *http.Request) string {
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	if ip := r.Header.Get("X-Real-Ip"); ip != "" {
		return ip
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", clientIP(r), r.Method, r.URL.RequestURI())

	headers := make(map[string]string, len(r.Header))
	for k, v := range r.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	body, _ := io.ReadAll(r.Body)

	resp := echoResponse{
		Method:  r.Method,
		Path:    r.URL.RequestURI(),
		Headers: headers,
		Body:    string(body),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/{path...}", echoHandler)

	log.Printf("echo-server listening on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
