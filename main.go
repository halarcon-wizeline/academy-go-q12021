package main

import (
	"fmt"
	"log"
	// "errors"

	"github.com/labstack/echo"

	// "github.com/halarcon-wizeline/academy-go-q12021/infrastructure/datastore"
	// "github.com/halarcon-wizeline/academy-go-q12021/interface/controller"
	"github.com/halarcon-wizeline/academy-go-q12021/registry"

	// "fmt"
	// "log"
	// "net/http"
	// "encoding/json"
	// "encoding/csv"
	// "io"
	// "os"

)

func main() {

	r := registry.NewRegistry()

/*
	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	serverPort := "8081"
	fmt.Println("Server listen at http://localhost" + ":" + serverPort)
	if err := e.Start(":" + serverPort); err != nil {
		log.Fatalln(err)
	}
	*/
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