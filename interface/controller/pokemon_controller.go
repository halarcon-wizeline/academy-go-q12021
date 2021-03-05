package controller

import (
	"fmt"
	"net/http"
	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
	"github.com/halarcon-wizeline/academy-go-q12021/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemons(c Context) error
}

func NewPokemonController(us interactor.PokemonInteractor) PokemonController {
	return &pokemonController{us}
}

func (uc *pokemonController) GetPokemons(c Context) error {

	log.Fatalln("GetPokemons")

	var u []*model.Pokemon

	u, err := uc.pokemonInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}