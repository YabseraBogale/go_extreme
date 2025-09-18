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

	err = db.AutoMigrate(&actors.EmployeeContact{}, &actors.Employee{},
		&actors.Item{}, &actors.TransactionType{},
		&actors.ItemLog{}, &actors.Checkout{})
	if err != nil {
		log.Println(err)
	}
}
