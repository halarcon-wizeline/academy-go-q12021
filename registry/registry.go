package registry

import (
	// "github.com/jinzhu/gorm"
	"github.com/halarcon-wizeline/academy-go-q12021/interface/controller"
)

type registry struct {
	// db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

// func NewRegistry(db *gorm.DB) Registry {
// 	return &registry{db}
// }
func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}