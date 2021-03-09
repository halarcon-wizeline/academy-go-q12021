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
	Get() ([]model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (pokemonInteractor *pokemonInteractor) Get() ([]model.Pokemon, error) {

	log.Println("Get")

	pokemon, err := pokemonInteractor.PokemonRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return pokemonInteractor.PokemonPresenter.ResponsePokemons(pokemon), nil
}