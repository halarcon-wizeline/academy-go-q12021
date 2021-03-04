package model

type Pokemon struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
}

func (Pokemon) TableName() string { return "pokemons" }