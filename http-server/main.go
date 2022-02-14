package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8001"

	http.HandleFunc("/", getData)

	fmt.Println("Listening on localhost" + port)
	http.ListenAndServe(port, nil)
}

func getData(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("result"))
}