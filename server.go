package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func loadClientCAs() *x509.CertPool {
    clientCAs := x509.NewCertPool()

    caCert, err := os.ReadFile("cert.pem")
    if err != nil {
        log.Fatalf("Error reading CA cert: %v", err)
    }

    clientCAs.AppendCertsFromPEM(caCert)
    return clientCAs
}

func main() {

	port := ":3000"

	// handling routes
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintln(w, "Handling incoming orders.")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintln(w, "Handling users.")
	})

	// Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
        ClientAuth: tls.RequireAndVerifyClientCert, // enforce mutual TLS (mTLS)
        ClientCAs:  loadClientCAs(), 
	}

	// Create a custom server
	server := &http.Server{
		Addr:      port,
		Handler:   nil, // use default mux
		TLSConfig: tlsConfig,
	}

	// Enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server Listening on port", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

	// HTTP 1.1 server without TLS
	// err := http.ListenAndServe(port, nil)

	// if err != nil {
	//     log.Fatal("Error starting server:", err)
	// }

}

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("Received request with HTTP Version:", httpVersion)

	if r.TLS != nil {
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("Received request with TLS Version:", tlsVersion)
	} else {
		fmt.Println("Received request without TLS")
	}
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown TLS Version"
	}
}
