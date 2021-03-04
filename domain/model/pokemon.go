package model

type Pokemon struct {
	ID        int      `json:"id"`
	Name      string    `json:"name"`
}

func (Pokemon) TableName() string { return "pokemons" }