package main

import (
	// "fmt"
	// "log"
	// "errors"

	// "github.com/labstack/echo"

	// "github.com/halarcon-wizeline/academy-go-q12021/infrastructure/datastore"
	// "github.com/halarcon-wizeline/academy-go-q12021/interface/controller"
	// "github.com/halarcon-wizeline/academy-go-q12021/registry"

	// "fmt"
	// "log"
	// "net/http"
	// "encoding/json"
	// "encoding/csv"
	// "io"
	// "os"

  // "github.com/halarcon-wizeline/academy-go-q12021/domain/model"
  // "github.com/halarcon-wizeline/academy-go-q12021/usecase/interactor"
  // "github.com/halarcon-wizeline/academy-go-q12021/usecase/presenter"
  // "github.com/halarcon-wizeline/academy-go-q12021/usecase/repository"

  ip "github.com/halarcon-wizeline/academy-go-q12021/interface/presenter"
  // ir "github.com/halarcon-wizeline/academy-go-q12021/interface/repository"

  // "github.com/halarcon-wizeline/academy-go-q12021/interface/controller"
)

// type pokemonRepository struct {
// 	// db *gorm.DB
// }

// func (ur *pokemonRepository) FindAll(u []*model.Pokemon) ([]*model.Pokemon, error) {
// 	// var Err error 
// 	var ErrNotFound = errors.New("not found")
// 	ErrNotFound = nil
// 	// err := nil

// 	if ErrNotFound != nil {
// 		return nil, ErrNotFound
// 	}

// 	return u, nil
// }

/*
type pokemonPresenter struct {
}

func (up *pokemonPresenter) ResponsePokemons(us []*model.Pokemon) []*model.Pokemon {
	for _, u := range us {
		u.Name = "Po: " + u.Name
	}
	return us
}


func main() {
*/


	ip.NewPokemonPresenter()
	// ir.NewPokemonRepository()


	// p := model.Pokemon {1, "test"}
	// type pokemonInteractor struct {
	// 	PokemonRepository repository.PokemonRepository
	// 	PokemonPresenter  presenter.PokemonPresenter
	// }


	// fmt.Println(p.ID)
	// fmt.Println(p.Name)

	// db := datastore.NewDB()
	// db.LogMode(true)
	// defer db.Close()

	// r := registry.NewRegistry(db)

	// e := echo.New()
	// e = router.NewRouter(e, r.NewAppController())

	// fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	// if err := e.Start(":" + config.C.Server.Address); err != nil {
	// 	log.Fatalln(err)
	// }
}


/*
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	// articles := Articles {
	// 	Article {Title:"Test title", Desc:"Test description", Content:"Hello World"},
	// }

	articles := readCsvArticles("./infrastructure/datastore/articles.csv")

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func readCsvArticles(file string) []Article {
		// Open the file
	csvfile, err := os.Open(file)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	var articles []Article
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Article: %s %s %s\n", record[0], record[1], record[2])
		article := Article {Title:record[0], Desc:record[1], Content:record[2]}
		articles = append(articles, article)
	}
	return articles
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}*/