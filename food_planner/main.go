package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("/public/index.html")
		if err != nil {
			log.Fatalln(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			log.Fatalln(err)
		}

	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
