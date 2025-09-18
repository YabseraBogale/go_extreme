package actors

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeContact struct {
	gorm.Model
	EmergencyContactId int `gorm:"primaryKey"`
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
	EmployeeId         int `gorm:"primaryKey"`
	EmergencyContactId int `gorm:"foreignKey"`
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
	EmployeeId          int `gorm:"foreignKey"`
	ItemId              int `gorm:"primaryKey"`
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
	CreatedByEmployeeId int `gorm:"foreignKey"`
	UpdateByEmployeeId  int `gorm:"foreignKey"`
}

type TransactionType struct {
	gorm.Model
	TransactionTypeId int
	TypeName          string
}

type ItemLog struct {
	gorm.Model
	LogId           int    `gorm:"primaryKey"`
	ItemId          int    `gorm:"foreignKey"`
	TransactionName string `gorm:"foreignKey"`
	EmployeeId      int    `gorm:"foreignKey"`
	QuantityChanged int
	TransactionDate time.Time
	Description     string
}

type Checkout struct {
	gorm.Model
	CheckoutId   int `gorm:"primaryKey"`
	ItemId       int `gorm:"foreignKey"`
	EmployeeId   int `gorm:"foreignKey"`
	CheckoutDate time.Time
	ReturnDate   time.Time
	Notes        string
}
