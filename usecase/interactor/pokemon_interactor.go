package interactor

import (
	"log"
	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
	"github.com/halarcon-wizeline/academy-go-q12021/usecase/presenter"
	"github.com/halarcon-wizeline/academy-go-q12021/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	GetPokemons() ([]model.Pokemon, error)
	GetPokemon(pokemonId int) (model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (pokemonInteractor *pokemonInteractor) GetPokemons() ([]model.Pokemon, error) {

	log.Println("GetPokemons")

	pokemons, err := pokemonInteractor.PokemonRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return pokemonInteractor.PokemonPresenter.ResponsePokemons(pokemons), nil
}

func (pokemonInteractor *pokemonInteractor) GetPokemon(pokemonId int) (model.Pokemon, error) {

	log.Println("GetPokemon")
	var newPokemon model.Pokemon

	newPokemon, err := pokemonInteractor.PokemonRepository.Find(pokemonId)
	if err != nil {
		return newPokemon, err
	}

	return pokemonInteractor.PokemonPresenter.ResponsePokemon(newPokemon), nil
}