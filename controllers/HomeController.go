package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	currencies, err := JSONHandler()
	if err != nil {
		panic(err)
	}


	view, _ := template.ParseFiles("views/index.html")

	if err := view.Execute(w, map[string]interface{}{
		"currencies": currencies}); err != nil {
		log.Printf("%v", err)
		http.Error(w, "Something went wrong...", http.StatusInternalServerError)
	}
	return

}
