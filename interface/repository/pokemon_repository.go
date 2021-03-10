package repository

import (
	"errors"
	"log"
	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
	"github.com/halarcon-wizeline/academy-go-q12021/infrastructure/datastore"
)

type pokemonRepository struct {
}

type PokemonRepository interface {
	FindAll() ([]model.Pokemon, error)
	Find(pokemonId int) (model.Pokemon, error)
}

func NewPokemonRepository() PokemonRepository {
	return &pokemonRepository{}
}

func (pokemonRepository *pokemonRepository) FindAll() ([]model.Pokemon, error) {

	log.Println("FindAll")

	pokemons := datastore.NewPokemonDB()

	// if err != nil {
	// 	return nil, err
	// }

	return pokemons, nil
}

func (pokemonRepository *pokemonRepository) Find(pokemonId int) (model.Pokemon, error) {

	log.Println("Find")
	var newPokemon model.Pokemon

	pokemons := datastore.NewPokemonDB()

	// if err != nil {
	// 	return nil, err
	// }

	for _, pokemon := range pokemons {
		if pokemon.ID == pokemonId {
			newPokemon = pokemon
			return newPokemon, nil
		}
	}

	return newPokemon, errors.New("Pokemon NOT found")
}