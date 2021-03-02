package model

type Pokemon struct {
	ID        uint      `json:"ID"`
	Title     string    `json:"Name"`
}

func (Pokemon) TableName() string { return "pokemons" }