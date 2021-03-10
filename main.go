package main

import (

	"fmt"
	"log"
	"github.com/labstack/echo"
	"github.com/halarcon-wizeline/academy-go-q12021/infrastructure/router"
	"github.com/halarcon-wizeline/academy-go-q12021/registry"
	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
	"github.com/halarcon-wizeline/academy-go-q12021/infrastructure/datastore"

	// TODO mux approach
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

var pokemons []model.Pokemon
var serverPort string = "8081"

func main() {

	useCleanArchitecture := false

	if useCleanArchitecture == true {
		r := registry.NewRegistry()

		e := echo.New()
		e.Debug = true
		
		e = router.NewRouter(e, r.NewAppController())

		fmt.Println("Server listen at http://localhost" + ":" + serverPort)
		if err := e.Start(":" + serverPort); err != nil {
			log.Fatalln(err)
		}

	} else {
		// mux approach
		pokemonsDB, _ := datastore.NewPokemonDB()
		pokemons = pokemonsDB

		handleRequests()
	}
}

// mux GET /pokemons endpoint
func allPokemons(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: All Pokemons Endpoint")
	json.NewEncoder(w).Encode(pokemons)
}

// mux GET /pokemons/{id} endpoint
func catchPokemon(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL)

	fmt.Println("Endpoint Hit: Catch Pokemon Endpoint")

	params := mux.Vars(r)
	for _, pokemon := range pokemons {
		val, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatalln("Couldn't convert value", err)
		}

		if pokemon.ID == val {
			json.NewEncoder(w).Encode(pokemon)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Pokemon{})
}

// mux GET / endpoint
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

// mux routes
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	fmt.Println("Server listen at http://localhost" + ":" + serverPort)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/pokemons", allPokemons).Methods("GET")
	myRouter.HandleFunc("/pokemons/{id}", catchPokemon).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+serverPort, myRouter))
}