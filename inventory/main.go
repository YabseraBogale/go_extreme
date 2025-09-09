package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/dashboard/item/insert", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/dashboard/item/update/<id>", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/dashboard/item/delete/<id>", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/dashboard/item/select/<id>", func(w http.ResponseWriter, r *http.Request) {

	})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println(err)
	}

}
