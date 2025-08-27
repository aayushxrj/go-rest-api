package main

import (
	"fmt"
	"net/http"
)

func main(){

    // handling routes 
    http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(w, "Handling incoming orders." )
    })

    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(w, "Handling incoming users." )
    })


    port := ":3000"
    fmt.Println("Server Listening on port", port)
    http.ListenAndServe(port, nil)
}
