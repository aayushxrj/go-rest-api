package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main(){

    port := ":3000"

    // handling routes 
    http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Handling incoming orders." )
    })

    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Handling users." )
    })

    // Load the TLS cert and key 
    cert := "cert.pem"
    key := "key.pem"

    // Configure TLS 
    tlsConfig := &tls.Config{
        MinVersion: tls.VersionTLS12,
    }

    // Create a custom server 
    server := &http.Server{
        Addr: port,
        Handler: nil, // use default mux
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
