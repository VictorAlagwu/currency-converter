package routes

import (
	"converter/controllers"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func Init() *mux.Router {
	route := mux.NewRouter()
	//route.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./public"))))
	route.NotFoundHandler = http.HandlerFunc(notFound)

	route.HandleFunc("/", controllers.Home)
	route.HandleFunc("/getRate", controllers.GetRate).Methods("POST")
	return route
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	view, err := template.ParseFiles("views/errors/404.html")

	if err != nil {
		log.Fatal(err)
	}

	_ = view.Execute(w, nil)
}

