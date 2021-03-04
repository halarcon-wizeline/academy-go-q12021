package controller

type AppController interface {
	Pokemon interface{ PokemonController }
}