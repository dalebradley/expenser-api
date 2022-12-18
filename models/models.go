package models

import "time"

// Expense contains details of an expense
type ExpenseResourceDB struct {
	ID        string    `bson:"id"`
	Type      string    `bson:"type"`
	Amount    float64   `bson:"amount"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
	CreatedBy CreatedBy `bson:"created_by" json:"created_by"`
}

// CreatedByDB is the user who is creating the payment session
type CreatedBy struct {
	Email    string `bson:"email" json:"email"`
	Forename string `bson:"forename" json:"forename"`
	ID       string `bson:"id" json:"id"`
	Surname  string `bson:"surname" json:"surname"`
}

// Expense contains details of an expense
type ExpenseResourceRest struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CreatedBy CreatedBy `json:"created_by"`
}

// IncomingCreateExpenseRequest represents the incoming request when creating an expense
type IncomingCreateExpenseRequest struct {
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
}

// Expenses is a list of Expense items
type ExpensesRest struct {
	Expenses []ExpenseResourceRest `json:"expenses"`
}
