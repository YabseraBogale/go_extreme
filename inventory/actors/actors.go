package actors

import "time"

type Employee struct {
	EmployeeId  int
	FirstName   string
	MiddleName  string
	LastName    string
	PhoneNumber string
	Email       string
}

type Item struct {
	EmployeeId      int
	ItemId          int
	ItemName        string
	ItemDescription string
	ItemPrice       string
	ItemQuantity    string
	Categories      string
	Warehouse       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
