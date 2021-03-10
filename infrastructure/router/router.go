package router

import (
	"fmt"
	"log"
	"github.com/halarcon-wizeline/academy-go-q12021/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, appController controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pokemons",
		func(context echo.Context) error {
			log.Println("**********")
			log.Println("GET /pokemons")
			fmt.Println(fmt.Sprintf("%#v", context))
			return appController.GetPokemons(context)
	})

	e.GET("/pokemons/:id",
		func(context echo.Context) error {
			log.Println("**********")
			log.Println("GET /pokemons/:id")
			pokemonId := context.Param("id")
			log.Println(pokemonId)
			return appController.GetPokemon(context)
	})

	return e
}