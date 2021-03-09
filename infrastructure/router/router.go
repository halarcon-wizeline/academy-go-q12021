package router

import (
	"fmt"
	"net/http"
	"github.com/halarcon-wizeline/academy-go-q12021/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	fmt.Printf("NewRouter")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pokemons",
		func(context echo.Context) error {
			fmt.Printf("pokemons")
			return c.GetPokemons(context)
	})
	e.GET("/pokemons/:id",
		func(context echo.Context) error {
			fmt.Printf("pokemons/:id")
			id := context.Param("id")
			return context.String(http.StatusOK, id)
	})

	e.GET("/",
		func(context echo.Context) error {
			fmt.Printf("pokemons")
			return context.String(http.StatusOK, "Hello, World!")
	})

	return e
}