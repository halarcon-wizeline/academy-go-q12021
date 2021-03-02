package repository

import (
	// "github.com/jinzhu/gorm"
	"../domain/model"
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
	// err := ur.db.Find(&u).Error
	err := ""

	if err != nil {
		return nil, err
	}

	return u, nil
}