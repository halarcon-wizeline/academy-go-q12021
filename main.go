package main

import (

	// TODO Clean architecture imports
	// "net/http"
	"fmt"
	"log"
	"github.com/labstack/echo"
	"github.com/halarcon-wizeline/academy-go-q12021/infrastructure/router"
	"github.com/halarcon-wizeline/academy-go-q12021/registry"
	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
	"github.com/halarcon-wizeline/academy-go-q12021/infrastructure/datastore"
	// "github.com/labstack/echo/middleware"
	// "github.com/halarcon-wizeline/academy-go-q12021/interface/controller"

/*
	// TODO mux approach
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

var pokemons []model.Pokemon
var serverPort string = "8081"
/*

// GET /pokemons endpoint
func allPokemons(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: All Pokemons Endpoint")
	json.NewEncoder(w).Encode(pokemons)
}

// GET /pokemons/{id} endpoint
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

// GET / endpoint
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

// Read csv file
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
		pokemon := model.Pokemon {ID:id, Name:record[1]}
		pokemons = append(pokemons, pokemon)
	}
	return pokemons
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	fmt.Println("Server listen at http://localhost" + ":" + serverPort)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/pokemons", allPokemons).Methods("GET")
	myRouter.HandleFunc("/pokemons/{id}", catchPokemon).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+serverPort, myRouter))
}
*/
func main() {
	
	pokemonDB := datastore.NewPokemonDB()
	// log.Printf(pokemonDB[0].Name)

	// TODO add Clean architecture
	r := registry.NewRegistry(pokemonDB)

	e := echo.New()
	e.Debug = true


	// &controller.pokemonController{pokemonInteractor:(*interactor.pokemonInteractor)(0xc000066240)}
	e = router.NewRouter(e, r.NewAppController())
	// fmt.Println(fmt.Sprintf("%#v", e))

	/*
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	log.Println("**********")
	log.Println("NewRouter Starting")
	e.GET("/pokemons/:id",
		func(context echo.Context) error {
			log.Println("**********")
			log.Println("pokemons/:id")
			id := context.Param("id")
			return context.String(http.StatusOK, id)
	})

	e.GET("/hello",
		func(context echo.Context) error {
			log.Println("**********")
			log.Println("hello")
			return context.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
	*/

	fmt.Println("Server listen at http://localhost" + ":" + serverPort)
	if err := e.Start(":" + serverPort); err != nil {
		log.Fatalln(err)
	}



	// mux approach
	// pokemons = readCsvPokemons("./infrastructure/datastore/pokemons.csv")
	// handleRequests()
}