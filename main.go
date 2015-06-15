package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// router.GET("/cafes/location/:location", Index)
	router.GET("/cafes", Index)
	router.GET("/cafes/:id", Show)

	// log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Server starting at port 8080.")
	http.ListenAndServe(":8080", router)
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")

	query := r.URL.Query()
	if len(query["location"]) > 0 {
		lq := query["location"][0]
		fmt.Fprint(w, "Your location query: ", lq)
	}
}

func Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("id"))
}
