package service

import (
  "fmt"
  "net/http"
  "strconv"
  "sync"
  "time"
  "errors"

  "github.com/sirupsen/logrus"
  "github.com/halarcon-wizeline/academy-go-q12021/domain"
  "github.com/halarcon-wizeline/academy-go-q12021/infrastructure/datastore"
  "github.com/halarcon-wizeline/academy-go-q12021/infrastructure/externalapi"
)

// Service
type Service struct {
  logger  *logrus.Logger
}

// New creates a new Service layer
func New(
  logger  *logrus.Logger,
  ) (*Service, error) {
  return &Service{logger}, nil
}

// GetExternalPokemons logic
func (s *Service) GetExternalPokemons() (string, error) {
  fmt.Println("[controller] GetExternalPokemons")

  logger := s.logger.WithFields(logrus.Fields{"func": "Get Pokemons Api", })
  logger.Debug("in")

  pokemons, err := externalapi.GetPokemons()
  if err != nil {
    logger.WithError(err).Error("Getting external api")
    return "", errors.New("Can't find api")
  }

  file, err := datastore.CreatePokemonDB("./infrastructure/datastore/pokemons_api.csv", pokemons)
  if err != nil {
    logger.WithError(err).Error("Exporting csv file")
    return "", errors.New("Export failed")
  }
  fmt.Println(file)

  return file, nil
}

// GetLocalPokemons logic
func (s *Service) GetLocalPokemons() ([]domain.Pokemon, error) {
  fmt.Println("[controller] GetLocalPokemons")

  logger := s.logger.WithFields(logrus.Fields{"func": "Get Local Pokemons", })
  logger.Debug("in")

  pokemons, err := datastore.GetPokemonDB()
  if err != nil {
    logger.WithError(err).Error("Getting local pokemons")
    return []domain.Pokemon{}, errors.New("Pokemons db not found")
  }

  return pokemons, nil
}

// GetLocalPokemon logic
func (s *Service) GetLocalPokemon(id string) (domain.Pokemon, error) {

  var pokemon = domain.Pokemon{}
  fmt.Println("[controller] GetLocalPokemon")

  logger := s.logger.WithFields(logrus.Fields{"func": "Get Local Pokemon", })
  logger.Debug("in")

  pokemons, err := datastore.GetPokemonDB()
  if err != nil {
    logger.WithError(err).Error("Getting local pokemon")
    return domain.Pokemon{}, errors.New("Pokemons db not found")
  }

  val, err := strconv.Atoi(id)
  if err != nil {
    logger.WithError(err).Error("Couldn't convert value")
    return pokemon, errors.New("Couldn't convert value")
  }

  for _, poke := range pokemons {
    if poke.ID == val {
      return poke, nil
    }
  }
  return pokemon, nil
}

// GetLocalPokemonWorkers logic
func (s *Service) GetLocalPokemonWorkers(r *http.Request) ([]domain.Pokemon, error) {
  fmt.Println("[controller] GetLocalPokemonWorkers")

  var pokemons = []domain.Pokemon{}

  // Default values
  pItems := 10
  pItemsPerWorker := 2
  pWorkers := 5
  // Default to odd
  pType := 1

  // Retrieve params and overwrite values
  pType           = getQueryParams(pType, r, "type")
  pItems          = getQueryParams(pItems, r, "items")
  pItemsPerWorker = getQueryParams(pItemsPerWorker, r, "items_per_workers")
  pWorkers        = getQueryParams(pWorkers, r, "workers")


  logger := s.logger.WithFields(logrus.Fields{"func": "Get Local Pokemon Workers", "param:type": pType, "param:items": pItems, "param:items_per_workers": pItemsPerWorker, } )
  logger.Debug("in")

  // Retrieve pokemon csv file
  pokemonDB, err := datastore.GetPokemonDB()
  if err != nil {
    logger.WithError(err).Error("Getting local pokemon workers")
    return pokemons, errors.New("Pokemons db not found")
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

  pokemons = getPokemonsCaught(chanPokCaught)
  fmt.Printf("Pokemons caught: %v\n", pokemons)

  return pokemons, nil
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