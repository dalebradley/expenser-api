package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dalebradley/expenser-api/models"
	"github.com/dalebradley/expenser-api/services"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func HandleGetExpenses(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	spew.Dump(vars)
}

func HandleCreateExpense(w http.ResponseWriter, r *http.Request) {
	requestDecoder := json.NewDecoder(r.Body)
	var incomingCreateExpenseRequest models.IncomingCreateExpenseRequest
	err := requestDecoder.Decode(&incomingCreateExpenseRequest)
	if err != nil {
		err = fmt.Errorf("error decoding incoming request when creating expense: %s", err)
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	createdBy := models.CreatedBy{
		Email:    "deebrad@hotmail.co.uk",
		Forename: "Dale",
		Surname:  "Bradley",
		ID:       "1",
	}

	expenseResourseRest := models.ExpenseResourceRest{
		ID:        uuid.New().String(),
		Type:      incomingCreateExpenseRequest.Type,
		Amount:    incomingCreateExpenseRequest.Amount,
		CreatedAt: time.Now(),
		CreatedBy: createdBy,
	}
	_, status, err := services.CreateExpense(r, expenseResourseRest)
	if status == 400 {
		err = fmt.Errorf("error creating expense: %s", err)
		fmt.Println(err)
		w.WriteHeader((http.StatusBadRequest))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func HandleGetExpense(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	expenseResoureRest, status, err := services.GetExpense(r, vars["id"])
	if status == 400 {
		err = fmt.Errorf("error getting expense <%s> with err: %s", vars["id"], err)
		fmt.Println(err)
		w.WriteHeader((http.StatusBadRequest))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(expenseResoureRest)
	if err != nil {
		err = fmt.Errorf("error writing response for GET expense: %s", err)
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
