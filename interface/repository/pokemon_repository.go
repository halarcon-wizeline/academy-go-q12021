package repository

import (
	"log"
	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
	"github.com/halarcon-wizeline/academy-go-q12021/infrastructure/datastore"
)

type pokemonRepository struct {
}

type PokemonRepository interface {
	FindAll() ([]model.Pokemon, error)
}

func NewPokemonRepository() PokemonRepository {
	return &pokemonRepository{}
}

func (pokemonRepository *pokemonRepository) FindAll() ([]model.Pokemon, error) {

	log.Println("FindAll")

	pokemons := datastore.NewPokemonDB()

	return pokemons, nil
}
