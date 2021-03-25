package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Controller
type Controller interface {
  GetExternalPokemons(w http.ResponseWriter, r *http.Request)
  GetLocalPokemons(w http.ResponseWriter, r *http.Request)
  GetLocalPokemon(w http.ResponseWriter, r *http.Request)
  GetLocalPokemonWorkers(w http.ResponseWriter, r *http.Request)
}

// New returns router instance which is used in main package to register handlers.
func New(controller Controller) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homeLink)
	r.HandleFunc("/pokemons_api", controller.GetExternalPokemons).Methods(http.MethodGet).Name("pokemons_api")
	
	r.HandleFunc("/pokemons/{id}", controller.GetLocalPokemon).Methods(http.MethodGet).Name("pokemon")
	r.HandleFunc("/pokemons", controller.GetLocalPokemons).Methods(http.MethodGet).Name("pokemons")

	// http://localhost:8080/pokemon_workers?type=even&items=10&items_per_workers=3&workers=3
	r.HandleFunc("/pokemon_workers", controller.GetLocalPokemonWorkers).Methods(http.MethodGet).Name("pokemon_workers")

	return r
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
