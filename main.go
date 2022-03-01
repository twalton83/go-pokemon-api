package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"io"
)

func main() {
  fmt.Println("Starting server...")

	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler).
		Methods("GET")
	r.HandleFunc("/pokemon/{name}", PokemonGetHandler).
		Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO WORLD!"))
}

func PokemonGetHandler(w http.ResponseWriter, r *http.Request) {
	// take POKEMON from query param
	pokemon := mux.Vars(r)["name"]
	// fetch PokeApi
	fmt.Println(pokemon)
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + pokemon)

	if err != nil {
		log.Fatalln(err)
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		} else {
			sb := string(body)
			w.Write([]byte(sb))
		}
	}


}
	


