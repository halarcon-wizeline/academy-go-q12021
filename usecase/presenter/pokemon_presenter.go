package presenter

import "../domain/model"

type PokemonPresenter interface {
	ResponsePokemons(u []*model.Pokemon) []*model.Pokemon
}