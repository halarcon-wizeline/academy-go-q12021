package repository

import (
	"errors"
	// "github.com/jinzhu/gorm"
	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
)

type pokemonRepository struct {
	// db *gorm.DB
}

type PokemonRepository interface {
	FindAll(u []*model.Pokemon) ([]*model.Pokemon, error)
}

// func NewPokemonRepository(db *gorm.DB) PokemonRepository {
func NewPokemonRepository() PokemonRepository {
	// return &pokemonRepository{db}
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