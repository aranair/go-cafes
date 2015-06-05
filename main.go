package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/cafes", Index)
	router.GET("/cafes/:id", Show)

	// log.Fatal(http.ListenAndServe(":8080", router))
	http.ListenAndServe(":8080", router)
	fmt.Println("Server started at port 8080.")
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
	fmt.Println(ps)
}

func Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("id"))
}
