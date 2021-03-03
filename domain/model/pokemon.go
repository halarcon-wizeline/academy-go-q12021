package model

type Pokemon struct {
	ID        uint      `json:"ID"`
	Name      string    `json:"Name"`
}

func (Pokemon) TableName() string { return "pokemons" }