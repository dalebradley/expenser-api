package routes

import (
	"net/http"

	"github.com/dalebradley/expenser-api/handlers"
	"github.com/gorilla/mux"
)

func Register(mainRouter *mux.Router) {
	r := mainRouter.PathPrefix("/expenses").Subrouter()
	r.HandleFunc("/create", handlers.ExpensesHandler).Methods("POST")
	http.Handle("/", r)
}
