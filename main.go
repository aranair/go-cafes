package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/cafes", Index)
	router.GET("/cafes/:id", Show)

	// log.Fatal(http.ListenAndServe(":8080", router))
	http.ListenAndServe(":8080", router)
}
