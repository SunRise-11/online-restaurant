package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func frontpageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/homepage.html"))
	tmpl.Execute(w, tmpl)
}
func orderHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/order.html"))
	tmpl.Execute(w, tmpl)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/home/", frontpageHandler).Methods("GET")
	r.HandleFunc("/order/", orderHandler).Methods("GET")
	r.HandleFunc("/meal", getMealHandler).Methods("GET")

	//serve static files
	staticFileDirectory := http.Dir("./static")
	fs := http.FileServer(staticFileDirectory)
	staticFileHandler := http.StripPrefix("/static/", fs)
	http.Handle("/static/", staticFileHandler)
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	return r
}

func main() {

	fmt.Println("Starting server...")
	// insert your postgres server username and password
	connString := "user=postgres password=12345 dbname=Foodie sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})
	// Declare a new router
	r := newRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
