package presenter

import "https://github.com/halarcon-wizeline/academy-go-q12021/domain/model"

type pokemonPresenter struct {
}

type PokemonPresenter interface {
	ResponsePokemons(us []*model.Pokemon) []*model.Pokemon
}

func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

func (up *pokemonPresenter) ResponsePokemons(us []*model.Pokemon) []*model.Pokemon {
	for _, u := range us {
		u.ID = "Pokemon " + u.Name
	}
	return us
}