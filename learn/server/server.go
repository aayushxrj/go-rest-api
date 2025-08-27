package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){

	// creating the handler function 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Hello, Server!")
	})

	// const serverAddr string = "127.0.0.1:3000"
	const port string = ":8080" // simplified

	// starting the server
	fmt.Println("Server Listening on port", port)
	err := http.ListenAndServe(port, nil) // address and handler (nil means default mux)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}


