package main

import (
	"fmt"
	"log"
	"github.com/labstack/echo"
	"github.com/halarcon-wizeline/academy-go-q12021/infrastructure/router"
	"github.com/halarcon-wizeline/academy-go-q12021/registry"
	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
)

var pokemons []model.Pokemon
var serverPort string = "8081"

func main() {

	r := registry.NewRegistry()

	e := echo.New()
	e.Debug = true
	
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + serverPort)
	if err := e.Start(":" + serverPort); err != nil {
		log.Fatalln(err)
	}
}