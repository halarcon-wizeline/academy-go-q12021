package repository

import (
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
	var ErrNotFound = errors.New("not found")
	ErrNotFound = nil

	if ErrNotFound != nil {
		return nil, ErrNotFound
	}

	return u, nil
}