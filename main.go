package main

import (
	"fmt"
	"log"

	//"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"

	// "../config"
	// "../infrastructure/datastore"
	"https://github.com/halarcon-wizeline/academy-go-q12021/infrastructure/router"
	// "../registry"
)

func main() {
	// config.ReadConfig()

	// db := datastore.NewDB()
	// db.LogMode(true)
	// defer db.Close()

	// r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}