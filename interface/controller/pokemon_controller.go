package controller

import (
	"log"
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"github.com/halarcon-wizeline/academy-go-q12021/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemons(context echo.Context) error
	GetPokemon(context echo.Context) error
}

func NewPokemonController(pokemonInteractor interactor.PokemonInteractor) PokemonController {
	return &pokemonController{pokemonInteractor}
}

func (pokemonController *pokemonController) GetPokemons(context echo.Context) error {
	log.Println("GetPokemons")

	pokemons, err := pokemonController.pokemonInteractor.GetPokemons()
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, pokemons)
}

func (pokemonController *pokemonController) GetPokemon(context echo.Context) error {
	log.Println("GetPokemon")

	pokemonId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	pokemon, err := pokemonController.pokemonInteractor.GetPokemon(pokemonId)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, pokemon)
}