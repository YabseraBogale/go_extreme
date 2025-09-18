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
	gorm.Model
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
	gorm.Model
	EmployeeId          int
	Employee            Employee
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
	gorm.Model
	TransactionTypeId int
	TypeName          string
}

type ItemLog struct {
	gorm.Model
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
	gorm.Model
	CheckoutId   int
	ItemId       int
	Item         Item
	EmployeeId   int
	Employee     Employee
	CheckoutDate time.Time
	ReturnDate   time.Time
	Notes        string
}
