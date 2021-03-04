package main

import (
	// "fmt"
	// "log"
	// "errors"

	// "github.com/labstack/echo"

	// "github.com/halarcon-wizeline/academy-go-q12021/infrastructure/datastore"
	// "github.com/halarcon-wizeline/academy-go-q12021/interface/controller"
	// "github.com/halarcon-wizeline/academy-go-q12021/registry"

	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"encoding/csv"
	"io"
	"os"

	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"

)

/*
func main() {

	// r := registry.NewRegistry()

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	serverPort := "8081"
	fmt.Println("Server listen at http://localhost" + ":" + serverPort)
	if err := e.Start(":" + serverPort); err != nil {
		log.Fatalln(err)
	}
}
	*/

type Pokemons []model.Pokemon

func allPokemons(w http.ResponseWriter, r *http.Request) {

	pokemons := readCsvPokemons("./infrastructure/datastore/pokemons.csv")

	fmt.Println("Endpoint Hit: All Pokemons Endpoint")
	json.NewEncoder(w).Encode(pokemons)
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
		fmt.Printf("Pokemon: %s %s\n", record[0], record[1])
		pokemon := model.Pokemon {ID:record[0], Name:record[1]}
		pokemons = append(pokemons, pokemon)
	}
	return pokemons
}

func handleRequests() {
	serverPort := "8081"
	fmt.Println("Server listen at http://localhost" + ":" + serverPort)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/pokemons", allPokemons)
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}

func main() {
	handleRequests()
}

