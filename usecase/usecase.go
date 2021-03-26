package usecase

import (
	"net/http"

	"github.com/halarcon-wizeline/academy-go-q12021/domain"
)

// UseCase struct
type UseCase struct {
	service Service
}

// Service interface
type Service interface {
  GetExternalPokemons() (file string, error error)
  GetLocalPokemons() ([]domain.Pokemon, error)
  GetLocalPokemon(id string) (domain.Pokemon, error)
  GetLocalPokemonWorkers(r *http.Request) ([]domain.Pokemon, error)
}

// New UseCase
func New(service Service) *UseCase {
	return &UseCase{service}
}

// GetExternalPokemons logic
func (u *UseCase) GetExternalPokemons() (string, error) {

	resp, err := u.service.GetExternalPokemons()
	return resp, err
}

// GetLocalPokemons logic
func (u *UseCase) GetLocalPokemons() ([]domain.Pokemon, error) {

	resp, err := u.service.GetLocalPokemons()
	return resp, err
}

// GetLocalPokemon logic
func (u *UseCase) GetLocalPokemon(id string) (domain.Pokemon, error) {

	resp, err := u.service.GetLocalPokemon(id)
	return resp, err
}

// GetLocalPokemonWorkers logic
func (u *UseCase) GetLocalPokemonWorkers(r *http.Request) ([]domain.Pokemon, error) {

	resp, err := u.service.GetLocalPokemonWorkers(r)
	return resp, err
}
