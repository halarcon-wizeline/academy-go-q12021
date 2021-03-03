package interactor

import (
	"https://github.com/halarcon-wizeline/academy-go-q12021/domain/model"
	"https://github.com/halarcon-wizeline/academy-go-q12021/usecase/presenter"
	"https://github.com/halarcon-wizeline/academy-go-q12021/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	Get(u []*model.Pokemon) ([]*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (us *pokemonInteractor) Get(u []*model.Pokemon) ([]*model.Pokemon, error) {
	u, err := us.PokemonRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.PokemonPresenter.ResponsePokemons(u), nil
}