package controller

import (
  "fmt"
  "net/http"
  "strconv"
  "github.com/gorilla/mux"

  "github.com/unrolled/render"
  "github.com/halarcon-wizeline/academy-go-q12021/domain"
)

// UseCase interface
type UseCase interface {
  GetExternalPokemons() (string, error)
  GetLocalPokemons() ([]domain.Pokemon, error)
  GetLocalPokemon(string) (domain.Pokemon, error)
  GetLocalPokemonWorkers(int, int, int, int) ([]domain.Pokemon, error)
}

// Controller struct
type Controller struct {
  useCase UseCase
  render  *render.Render
}

// New returns a controller
func New(
  u UseCase,
  r *render.Render,
) *Controller {
  return &Controller{u, r}
}

var pItems = 10
var pItemsPerWorker = 2
var pWorkers = 5
var pType = 1

func init() {
  // Default values
  pItems = 10
  pItemsPerWorker = 2
  pWorkers = 5
  // Default to odd
  pType = 1
}

// GetExternalPokemons logic
func (c *Controller) GetExternalPokemons(w http.ResponseWriter, r *http.Request) {
  fmt.Println("[controller] GetExternalPokemons")

  file, err := c.useCase.GetExternalPokemons()
  if err != nil {
    c.render.Text(w, http.StatusOK, "Error")
  }

  c.render.Text(w, http.StatusOK, file + " created")
}

// GetLocalPokemons logic
func (c *Controller) GetLocalPokemons(w http.ResponseWriter, r *http.Request) {
  fmt.Println("[controller] GetLocalPokemons")

  pokemons, err := c.useCase.GetLocalPokemons()
  if err != nil {
    c.render.Text(w, http.StatusOK, "Pokemons db not found")
    return 
  }

  w.Header().Set("Content-Type", "application/json")
  c.render.JSON(w, http.StatusOK, pokemons)
}

// GetLocalPokemon logic
func (c *Controller) GetLocalPokemon(w http.ResponseWriter, r *http.Request) {
  fmt.Println("[controller] GetLocalPokemon")

  params := mux.Vars(r)
  pokemon, err := c.useCase.GetLocalPokemon(params["id"])
  if err != nil {
    c.render.Text(w, http.StatusOK, "Pokemon not found")
    return 
  }

  w.Header().Set("Content-Type", "application/json")
  c.render.JSON(w, http.StatusOK, pokemon)
}

// GetLocalPokemonWorkers logic
func (c *Controller) GetLocalPokemonWorkers(w http.ResponseWriter, r *http.Request) {
  fmt.Println("[controller] GetLocalPokemonWorkers")

  // Retrieve params and overwrite values
  pType           = getQueryParams(pType, r, "type")
  pItems          = getQueryParams(pItems, r, "items")
  pItemsPerWorker = getQueryParams(pItemsPerWorker, r, "items_per_workers")
  pWorkers        = getQueryParams(pWorkers, r, "workers")

  pokemons, err := c.useCase.GetLocalPokemonWorkers(pType, pItems, pItemsPerWorker, pWorkers)
  if err != nil {
    c.render.Text(w, http.StatusOK, "Pokemon not found")
    return 
  }

  w.Header().Set("Content-Type", "application/json")
  c.render.JSON(w, http.StatusOK, pokemons)
}

func getQueryParams(defaultValue int, r *http.Request, index string) int {
  param := r.URL.Query().Get(index)
  newParam := defaultValue

  if index == "type" {
    switch param {
      case "odd":
        newParam = 1
      case "even":
        newParam = 2
    }
  } else {
    intParam, err := strconv.Atoi(param)
    if err == nil {
      newParam = intParam
    }
  }

  return newParam
}
