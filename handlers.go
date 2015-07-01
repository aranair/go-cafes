package gocafes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

type AppContext struct {
	db *sql.DB
}

// TODO: Add in db app context
func NewAppContext(db *sql.DB) AppContext {
	return AppContext{db: db}
}

func (ac *AppContext) IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")

	query := r.URL.Query()
	if len(query.Get("location")) > 0 {
		lq := query.Gkt("location")[0]
		fmt.Fprint(w, "Your location query: ", lq)
	}

	fmt.Println(context.Get(r, "params"))
	params := context.Get(r, "params").(httprouter.Params)
	fmt.Println(params.ByName("location"))
}

func (ac *AppContext) ShowHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	fmt.Fprintf(w, "hello, %s!\n", params.ByName("id"))
}
