package registry

import (
	"github.com/halarcon-wizeline/academy-go-q12021/interface/controller"
)

type registry struct {
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}