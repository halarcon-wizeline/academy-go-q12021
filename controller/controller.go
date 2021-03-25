package controller

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "strconv"
  "sync"
  "time"

  "github.com/sirupsen/logrus"
  "github.com/unrolled/render"
  "github.com/halarcon-wizeline/academy-go-q12021/domain"
  "github.com/halarcon-wizeline/academy-go-q12021/infrastructure/datastore"
  "github.com/halarcon-wizeline/academy-go-q12021/infrastructure/externalapi"
)

// UseCase interface
type UseCase interface {
}

// Controller struct
type Controller struct {
  useCase UseCase
  logger  *logrus.Logger
  render  *render.Render
}

// New returns a controller
func New(
  u UseCase,
  logger *logrus.Logger,
  r *render.Render,
) *Controller {
  return &Controller{u, logger, r}
}

// GetExternalPokemons logic
func (c *Controller) GetExternalPokemons(w http.ResponseWriter, r *http.Request) {
  fmt.Println("[controller] GetExternalPokemons")

  logger := c.logger.WithFields(logrus.Fields{"func": "Get Pokemons Api", })
  logger.Debug("in")

  pokemons, err := externalapi.GetPokemons()
  if err != nil {
    logger.WithError(err).Error("Getting external api")
    c.render.Text(w, 200, "Can't find api")
    return
  }

  file, err := datastore.CreatePokemonDB("./infrastructure/datastore/pokemons_api.csv", pokemons)
  if err != nil {
    logger.WithError(err).Error("Exporting csv file")
    c.render.Text(w, 200, "Export failed")
    return
  }
  fmt.Println(file)

  c.render.Text(w, 200, file + "created")
}

// GetLocalPokemons logic
func (c *Controller) GetLocalPokemons(w http.ResponseWriter, r *http.Request) {
  fmt.Println("[controller] GetLocalPokemons")

  logger := c.logger.WithFields(logrus.Fields{"func": "Get Local Pokemons", })
  logger.Debug("in")

  pokemons, err := datastore.GetPokemonDB()
  if err != nil {
    logger.WithError(err).Error("Getting local pokemons")
    c.render.Text(w, 200, "Pokemons db not found")
    return 
  }

  json.NewEncoder(w).Encode(pokemons)
}

// GetLocalPokemon logic
func (c *Controller) GetLocalPokemon(w http.ResponseWriter, r *http.Request) {
  fmt.Println("[controller] GetLocalPokemon")

  logger := c.logger.WithFields(logrus.Fields{"func": "Get Local Pokemon", })
  logger.Debug("in")

  pokemons, err := datastore.GetPokemonDB()
  if err != nil {
    logger.WithError(err).Error("Getting local pokemon")
    c.render.Text(w, 200, "Pokemons db not found")
    return 
  }

  params := mux.Vars(r)
  for _, pokemon := range pokemons {
    val, err := strconv.Atoi(params["id"])
    if err != nil {
      logger.WithError(err).Error("Couldn't convert value")
      c.render.Text(w, 200, "Couldn't convert value")
      return 
    }

    if pokemon.ID == val {
      json.NewEncoder(w).Encode(pokemon)
      return
    }
  }
  json.NewEncoder(w).Encode(&domain.Pokemon{})
}

// GetLocalPokemonWorkers logic
func (c *Controller) GetLocalPokemonWorkers(w http.ResponseWriter, r *http.Request) {
  fmt.Println("[controller] GetLocalPokemonWorkers")

  // Default values
  pItems := 10
  pItemsPerWorker := 2
  pWorkers := 5
  // Default to odd
  pType := 1

  // Retrieve params and overwrite values
  paramType := r.URL.Query().Get("type")
  switch paramType {
    case "odd":
      pType = 1
    case "even":
      pType = 2
  }
  paramItems, err := strconv.Atoi(r.URL.Query().Get("items"))
  if err == nil {
    pItems = paramItems
  }
  paramItemsPerWorker, err := strconv.Atoi(r.URL.Query().Get("items_per_workers"))
  if err == nil {
    pItemsPerWorker = paramItemsPerWorker
  }
  paramWorkers, err := strconv.Atoi(r.URL.Query().Get("workers"))
  if err == nil {
    pWorkers = paramWorkers
  }

  logger := c.logger.WithFields(logrus.Fields{"func": "Get Local Pokemon Workers", "param:type": paramType, "param:items": pItems, "param:items_per_workers": pItemsPerWorker, } )
  logger.Debug("in")

  // Retrieve pokemon csv file
  pokemonDB, err := datastore.GetPokemonDB()
  if err != nil {
    logger.WithError(err).Error("Getting local pokemon workers")
    c.render.Text(w, 200, "Pokemons db not found")
    return 
  }

  var wg sync.WaitGroup

  chanPokToCatch := make(chan int, pItems)
  chanPokCaught := make(chan domain.Pokemon, pItems)

  // Generate lists of pokemons to be captured
  for i := 0; i < pItems; i++ {
    chanPokToCatch <- lookForPokemon(i, pType)
  }
  close(chanPokToCatch)

  for i := 1; i <= pWorkers; i++ {
    fmt.Println("Main: Starting worker", i)
    wg.Add(1)
    go worker(&wg, pokemonDB, i, chanPokToCatch, chanPokCaught, pItemsPerWorker)
  }

  fmt.Println("Main: Waiting for workers to finish")
  wg.Wait()
  fmt.Println("Main: Completed")

  pokemons := getPokemonsCaught(chanPokCaught)
  fmt.Printf("Pokemons caught: %v\n", pokemons)

  json.NewEncoder(w).Encode(pokemons)
}

func worker(wg *sync.WaitGroup, pokemonDB []domain.Pokemon, workerId int, chanPokToCatch <- chan int, chanPokCaught chan <- domain.Pokemon, pItemsPerWorker int) {
  defer wg.Done()

  // worker pokemon caught counter
  found := 0
  for job := range chanPokToCatch {

    // Items per worker validation
    if found >= pItemsPerWorker {
      fmt.Printf("Worker %v stopped working, got %v pokemon(s)\n", workerId, found)
      break;
    }

    fmt.Printf("Worker %v looking for pokemon %v\n", workerId, job)
    // Items completed validation
    if job > len(pokemonDB)-1 {
      fmt.Printf("Pokemon %v does not exists\n", job)
      break;
    }

    chanPokCaught <- pokemonDB[job]
    time.Sleep(time.Second)
    fmt.Printf("Worker %v got pokemon %v\n", workerId, job)
    found = found + 1
  }
}

func lookForPokemon(index int, pType int) int {
  pokeNumber := index
  switch pType {
    case 1: pokeNumber = 2 * index + 1
    case 2: pokeNumber = 2 * index
  }
  return pokeNumber
}

func getPokemonsCaught(chanPokCaught <- chan domain.Pokemon) []domain.Pokemon {

  total := len(chanPokCaught)
  fmt.Printf("Total pokemons caught: %v\n", total)
  results := make([]domain.Pokemon, total)

  for j := 0; j < total; j++ {
    results[j] = <-chanPokCaught
  }

  return results
}