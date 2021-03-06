package presenter

import "github.com/halarcon-wizeline/academy-go-q12021/domain/model"

type PokemonPresenter interface {
	ResponsePokemons(pokemon []model.Pokemon) []model.Pokemon
	ResponsePokemon(pokemon model.Pokemon) model.Pokemon
}