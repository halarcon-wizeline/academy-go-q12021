package repository

import (
	"fmt"
	"log"
	"io"
	"os"
	"strconv"
	"errors"
	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
)

type pokemonRepository struct {
}

type PokemonRepository interface {
	FindAll(u []*model.Pokemon) ([]*model.Pokemon, error)
}

func NewPokemonRepository() PokemonRepository {
	return &pokemonRepository{}
}

func (ur *pokemonRepository) FindAll(u []*model.Pokemon) ([]*model.Pokemon, error) {

	fmt.Printf("FindAll")

	pokemons = readCsvPokemons("../infrastructure/datastore/pokemons.csv")

	// Open the file
	csvfile, err := os.Open(file)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
		return nil, err
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	// var pokemons []model.Pokemon
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
	return pokemons, nil
}
