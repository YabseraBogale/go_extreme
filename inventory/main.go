package main

import (
	"log"

	"github.com/YabseraBogale/go_extreme/inventory/actors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	var employee actors.Employee
	var item actors.Item
	var item_log actors.ItemLog
	var checkout actors.Checkout
	var transaction_type actors.TransactionType
	var emergency_contact actors.EmployeeContact

	db.AutoMigrate(&emergency_contact, &employee, &item, &item_log, &checkout, &transaction_type)

}
