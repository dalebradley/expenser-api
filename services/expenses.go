package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dalebradley/expenser-api/models"
	"github.com/dalebradley/expenser-api/storage"

	"honnef.co/go/tools/config"
)

// Expenseservice contains the Repository for db access
type ExpensesService struct {
	Repository storage.Mongo
	Config     config.Config
}

func CreateExpense(req *http.Request, createExpenseResource models.IncomingCreateExpenseRequest) (*models.ExpenseResourceRest, int, error) {
	// Add trace logs
	createdBy := models.CreatedBy{
		Email:    "deebrad@hotmail.co.uk",
		Forename: "Dale",
		Surname:  "Bradley",
		ID:       "1",
	}
	expenseResourceDB := models.ExpenseResourceDB{
		Type:      "test",
		ID:        "1",
		CreatedAt: time.Now(),
		CreatedBy: createdBy,
	}

	err := storage.CreateExpenseResource(expenseResourceDB)
	if err != nil {
		err = fmt.Errorf("error creating expense resource: [%v]", err)
		//LOG ERROR
		return nil, 500, err
	}
	expenseResourseRest := models.ExpenseResourceRest(expenseResourceDB)
	return &expenseResourseRest, 200, nil
}
