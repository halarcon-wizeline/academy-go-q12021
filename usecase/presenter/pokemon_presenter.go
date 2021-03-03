package presenter

import "https://github.com/halarcon-wizeline/academy-go-q12021/domain/model"

type PokemonPresenter interface {
	ResponsePokemons(u []*model.Pokemon) []*model.Pokemon
}