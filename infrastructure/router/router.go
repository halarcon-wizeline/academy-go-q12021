package router

import (
	"log"
	"net/http"
	"github.com/halarcon-wizeline/academy-go-q12021/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e echo.Echo, c controller.AppController) echo.Echo {
	log.Println("**********")
	log.Println("NewRouter Starting")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/newpokemons",
		func(context echo.Context) error {
			log.Println("**********")
			log.Println("pokemons")
			return c.GetPokemons(context)
	})
	e.GET("/pokemons/:id",
		func(context echo.Context) error {
			log.Println("**********")
			log.Println("pokemons/:id")
			id := context.Param("id")
			return context.String(http.StatusOK, id)
	})

	e.GET("/hello",
		func(context echo.Context) error {
			log.Println("**********")
			log.Println("hello")
			return context.String(http.StatusOK, "Hello, World!")
	})

	return e
}