package model

type Pokemon struct {
	ID        string      `json:"id"`
	Name      string    `json:"name"`
}

func (Pokemon) TableName() string { return "pokemons" }