package routes

import (
	"github.com/dalebradley/expenser-api/handlers"
	"github.com/gorilla/mux"
)

func Register(mainRouter *mux.Router) {
	r := mainRouter.PathPrefix("/expenses").Subrouter()
	r.HandleFunc("/", handlers.HandleGetExpenses).Methods("GET")
	r.HandleFunc("/create", handlers.HandleCreateExpense).Methods("POST")
}
