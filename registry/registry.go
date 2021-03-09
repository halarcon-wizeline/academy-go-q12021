package registry

import (
	"github.com/halarcon-wizeline/academy-go-q12021/domain/model"
	"github.com/halarcon-wizeline/academy-go-q12021/interface/controller"
)

type registry struct {
	pokemonDB []model.Pokemon
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(pokemonDB []model.Pokemon) Registry {
	return &registry{pokemonDB}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}