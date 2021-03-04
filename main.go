package main

import (
	"fmt"
	"log"
	// "errors"

	"github.com/labstack/echo"

	// "github.com/halarcon-wizeline/academy-go-q12021/infrastructure/datastore"
	// "github.com/halarcon-wizeline/academy-go-q12021/interface/controller"
	// "github.com/gorilla/mux"
	"github.com/halarcon-wizeline/academy-go-q12021/infrastructure/router"
	"github.com/halarcon-wizeline/academy-go-q12021/registry"

/*
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"github.com/gorilla/mux"

	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
*/
)


func main() {

	r := registry.NewRegistry()

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	serverPort := "8081"
	fmt.Println("Server listen at http://localhost" + ":" + serverPort)
	if err := e.Start(":" + serverPort); err != nil {
		log.Fatalln(err)
	}
}



/*
var pokemons []model.Pokemon

func allPokemons(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: All Pokemons Endpoint")
	json.NewEncoder(w).Encode(pokemons)
}

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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func readCsvPokemons(file string) []model.Pokemon {
	// Open the file
	csvfile, err := os.Open(file)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	var pokemons []model.Pokemon
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Reading pokemon: %s %s\n", record[0], record[1])
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalln("Error: Pokemon: %s does not have a valid ID\n", record[1])
		}
		var val int = id
		pokemon := model.Pokemon {ID:id, Name:record[1]}
		pokemons = append(pokemons, pokemon)
	}
	return pokemons
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	serverPort := "8081"
	fmt.Println("Server listen at http://localhost" + ":" + serverPort)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/pokemons", allPokemons).Methods("GET")
	myRouter.HandleFunc("/pokemons/{id}", catchPokemon).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+serverPort, myRouter))
}

func main() {

	pokemons = readCsvPokemons("./infrastructure/datastore/pokemons.csv")

	handleRequests()
}
*/