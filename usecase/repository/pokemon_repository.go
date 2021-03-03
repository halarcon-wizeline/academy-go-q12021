package repository

import "https://github.com/halarcon-wizeline/academy-go-q12021/domain/model"

type PokemonRepository interface {
	FindAll(u []*model.Pokemon) ([]*model.Pokemon, error)
}