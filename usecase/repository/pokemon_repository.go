package repository

import "github.com/halarcon-wizeline/academy-go-q12021/domain/model"

type PokemonRepository interface {
	FindAll() ([]model.Pokemon, error)
	Find(pokemonId int) (model.Pokemon, error)
}