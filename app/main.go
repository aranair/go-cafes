package main

import (
	"fmt"
	"net/http"

	gc "github.com/aranair/gocafes"
	router "github.com/aranair/gocafes/router"
	"github.com/justinas/alice"
)

func main() {
	ac := gc.NewAppContext()
	stack := alice.New()

	r := router.New()
	r.GET("/cafes", stack.ThenFunc(ac.IndexHandler))
	r.GET("/cafes/:id", stack.ThenFunc(ac.ShowHandler))

	http.ListenAndServe(":8080", r)
	fmt.Println("Server started at port 8080.")
}
