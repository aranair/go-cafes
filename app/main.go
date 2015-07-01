package main

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	gc "github.com/aranair/gocafes"
	router "github.com/aranair/gocafes/router"

	_ "github.com/lib/pq"

	"github.com/BurntSushi/toml"
	"github.com/justinas/alice"
)

type Config struct {
	DB database `toml:"database"`
}

type database struct {
	User     string
	Password string
}

type Cafe struct {
	Address string
	Name    string
	Score   float32
	Id      int
}

func main() {
	var conf Config
	if _, err := toml.DecodeFile("configs.toml", &conf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf)

	pqStr := "user=" + conf.DB.User + " password=" + conf.DB.Password + " dbname=gocafe sslmode=verify-full"
	db, err := sql.Open("postgres", pqStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// var cafe Cafe
	rows, err := db.Query("SELECT id FROM cafes WHERE id = $1", 1)
	fmt.Println(rows)
	// fmt.Println(cafe)

	ac := gc.NewAppContext(db)
	stack := alice.New()

	r := router.New()
	r.GET("/cafes", stack.ThenFunc(ac.IndexHandler))
	r.GET("/cafes/:id", stack.ThenFunc(ac.ShowHandler))

	fmt.Println("Server starting at port 8080.")
	http.ListenAndServe(":8080", r)
}
