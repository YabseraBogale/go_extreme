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
	Categories      Categories
	Warehouse       Warehouses
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Warehouses struct {
	WarehouseId int
	Name        string
	Location    string
}

type Categories struct {
	CategoryId int
	Name       string
}

type StockMovements struct {
	EmployeeId   int
	MovementId   int
	ItemId       Item
	WarehouseId  Warehouses
	ChangeQty    int
	MovementType string
	CreatedAt    time.Time
}
