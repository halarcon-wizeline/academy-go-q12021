package registry

import (
	"https://github.com/halarcon-wizeline/academy-go-q12021/interface/controller"
	ip "https://github.com/halarcon-wizeline/academy-go-q12021/interface/presenter"
	ir "https://github.com/halarcon-wizeline/academy-go-q12021/interface/repository"
	"https://github.com/halarcon-wizeline/academy-go-q12021/usecase/interactor"
	up "https://github.com/halarcon-wizeline/academy-go-q12021/usecase/presenter"
	ur "https://github.com/halarcon-wizeline/academy-go-q12021/usecase/repository"
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