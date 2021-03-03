package registry

import (
	// "github.com/halarcon-wizeline/academy-go-q12021/interface/controller"
	ip "github.com/halarcon-wizeline/academy-go-q12021/interface/presenter"
	ir "github.com/halarcon-wizeline/academy-go-q12021/interface/repository"
	"github.com/halarcon-wizeline/academy-go-q12021/usecase/interactor"
	up "github.com/halarcon-wizeline/academy-go-q12021/usecase/presenter"
	ur "github.com/halarcon-wizeline/academy-go-q12021/usecase/repository"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter())
}

func (r *registry) NewPokemonRepository() ur.PokemonRepository {
	return ir.NewPokemonRepository(r.db)
}

func (r *registry) NewPokemonPresenter() up.PokemonPresenter {
	return ip.NewPokemonPresenter()
}