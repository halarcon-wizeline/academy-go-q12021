package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"../interface/controller"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pokemons", func(context echo.Context) error { return c.GetPokemons(context) })

	return e
}