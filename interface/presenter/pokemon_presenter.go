package presenter

import (
	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
)

type pokemonPresenter struct {
}

type PokemonPresenter interface {
	ResponsePokemons(pokemonPresenter []model.Pokemon) []model.Pokemon
}

func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

func (pokemonPresenter *pokemonPresenter) ResponsePokemons(pokemons []model.Pokemon) []model.Pokemon {
	return pokemons
}