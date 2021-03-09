package datastore

import (
	"log"
	"fmt"
	"io"
	"os"
	"encoding/csv"
	"strconv"
	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
)

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

func NewPokemonDB() []model.Pokemon {

	var pokemons []model.Pokemon
  pokemons = readCsvPokemons("./infrastructure/datastore/pokemons.csv")

  return pokemons
}