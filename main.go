package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func main() {
  fmt.Println("Starting server...")

	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO WORLD!"))
}

