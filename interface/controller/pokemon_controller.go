package controller

import (
	"log"
	"net/http"
	"github.com/halarcon-wizeline/academy-go-q12021/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemons(context Context) error
}

func NewPokemonController(pokemonInteractor interactor.PokemonInteractor) PokemonController {
	return &pokemonController{pokemonInteractor}
}

func (pokemonController *pokemonController) GetPokemons(context Context) error {
	log.Println("GetPokemons")

	pokemon, err := pokemonController.pokemonInteractor.Get()

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, pokemon)
}