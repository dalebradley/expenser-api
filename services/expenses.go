package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dalebradley/expenser-api/models"
	"github.com/dalebradley/expenser-api/storage"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-playground/validator/v10"

	"honnef.co/go/tools/config"
)

// Expenseservice contains the Repository for db access
type ExpensesService struct {
	Repository storage.Mongo
	Config     config.Config
}

func CreateExpense(req *http.Request, createExpenseRequest models.ExpenseResourceRest) (*models.ExpenseResourceRest, int, error) {
	// Add trace logs
	// TODO: Hardcoded createdBy until auth implemented
	// createdBy := models.CreatedBy{
	// 	Email:    "deebrad@hotmail.co.uk",
	// 	Forename: "Dale",
	// 	Surname:  "Bradley",
	// 	ID:       "1",
	// }
	// expenseResourceDB := models.ExpenseResourceDB{
	// 	Type:      "test",
	// 	ID:        "1",
	// 	CreatedAt: time.Now(),
	// 	CreatedBy: createdBy,
	// }

	err := validateCreateExpenseRequestBody(createExpenseRequest)
	spew.Dump("AFTER VALIDATE")
	spew.Dump(err)
	if err != nil {
		spew.Dump("In ERROR")
		err = fmt.Errorf("invalid expense resource: [%v]", err)
		//LOG ERROR
		return nil, 400, err
	}

	err = storage.CreateExpenseResource(createExpenseRequest)
	if err != nil {
		err = fmt.Errorf("error creating expense resource: [%v]", err)
		//LOG ERROR
		return nil, 500, err
	}
	expenseResourseRest := models.ExpenseResourceRest(createExpenseRequest)
	return &expenseResourseRest, 200, nil
}

func GetExpense(req *http.Request, id string) (models.ExpenseResourceRest, int, error) {
	expenseResourceDB, err := storage.GetExpenseResource(id)
	if err != nil {
		//Handle
	}
	return models.ExpenseResourceRest(*expenseResourceDB), 200, nil
}

func validateCreateExpenseRequestBody(createExpenseRequest models.ExpenseResourceRest) error {
	validate := validator.New()
	spew.Dump("In VALIDATE")
	spew.Dump(createExpenseRequest)
	err := validate.Struct(createExpenseRequest)
	if err != nil {
		return err
	}
	// TODO: Check all fields
	if createExpenseRequest.ID == "" {
		return errors.New("empty ID")
	}
	if createExpenseRequest.Amount == 0 {
		return errors.New("no amount provided")
	}
	if createExpenseRequest.Type == "" {
		return errors.New("no type provided")
	}
	return nil
}
