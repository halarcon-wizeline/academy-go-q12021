package datastore

import (
	"log"
	"fmt"
	"io"
	"os"
	"errors"
	"encoding/csv"
	"strconv"
	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
)

// Read csv file
func readCsvPokemons(file string) ([]model.Pokemon, error) {

	var pokemons []model.Pokemon

	// Open the file
	csvfile, err := os.Open(file)
	if err != nil {
		return pokemons, errors.New("Couldn't open the csv file")
	}

	// Parse the file
	r := csv.NewReader(csvfile)

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
			return pokemons, errors.New("Error: Pokemon does not have a valid ID\n")
		}
		pokemon := model.Pokemon {ID:id, Name:record[1]}
		pokemons = append(pokemons, pokemon)
	}
	return pokemons, nil
}

func NewPokemonDB() ([]model.Pokemon, error) {

	pokemons, err := readCsvPokemons("./infrastructure/datastore/pokemons.csv")

	if err != nil {
		return nil, err
	}

	return pokemons, nil
}