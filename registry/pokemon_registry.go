package registry

import (
	"../interface/controller"
	ip "../interface/presenter"
	ir "../interface/repository"
	"../usecase/interactor"
	up "../usecase/presenter"
	ur "../usecase/repository"
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