package actors

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeContact struct {
	gorm.Model
	EmergencyContactId int
	FirstName          string
	MiddleName         string
	LastName           string
	PhoneNumber        string
	Email              string
	Location           string
	FyidaId            string
}

type Employee struct {
	EmployeeId         int
	EmergencyContactId int
	EmergencyContact   EmployeeContact
	FirstName          string
	MiddleName         string
	LastName           string
	PhoneNumber        string
	Email              string
	FyidaId            string
	Position           string
	Department         string
	TinNumber          int
	Location           string
	BankAccountNumber  int
	Salary             float32
}

type Item struct {
	Employee            Employee
	EmployeeId          int
	ItemId              int
	ItemName            string
	ItemDescription     string
	ItemPrice           string
	ItemQuantity        string
	Categories          string
	Location            string
	Unit                string
	Warehouse           string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	CreatedByEmployeeId int
	UpdateByEmployeeId  int
}

type TransactionType struct {
	TransactionTypeId int
	TypeName          string
}

type ItemLog struct {
	LogId           int
	ItemId          int
	Item            Item
	TransactionType TransactionType
	EmployeeId      int
	Employee        Employee
	QuantityChanged int
	TransactionDate time.Time
	Description     string
}

type Checkout struct {
	CheckoutId   int
	ItemId       int
	Item         Item
	EmployeeId   int
	Employee     Employee
	CheckoutDate time.Time
	ReturnDate   time.Time
	Notes        string
}
