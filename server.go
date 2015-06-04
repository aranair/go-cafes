package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("id"))
}

func main() {
	router := httprouter.New()
	router.GET("/cafes", Index)
	router.GET("/cafes/:id", Show)

	log.Fatal(http.ListenAndServe(":8080", router))
}
