package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dalebradley/expenser-api/models"
	"github.com/dalebradley/expenser-api/services"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

func HandleGetExpenses(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	spew.Dump(vars)
}

func HandleCreateExpense(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	requestDecoder := json.NewDecoder(r.Body)
	var incomingCreateExpenseRequest models.IncomingCreateExpenseRequest
	err := requestDecoder.Decode(&incomingCreateExpenseRequest)
	if err != nil {
		// Log error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	services.CreateExpense(r, incomingCreateExpenseRequest)
	w.WriteHeader(http.StatusOK)
	spew.Dump(vars)
}
