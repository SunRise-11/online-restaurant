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
func menuHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/menu.html"))
	tmpl.Execute(w, tmpl)
}
func orderHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/order.html"))
	tmpl.Execute(w, tmpl)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/home/", frontpageHandler).Methods("GET")
	r.HandleFunc("/menu/", menuHandler).Methods("GET")
	r.HandleFunc("/order/", orderHandler).Methods("GET")

	// database query functions
	r.HandleFunc("/", getMealHandler).Methods("GET")
	r.HandleFunc("/order", createOrderHandler).Methods("POST")
	r.HandleFunc("/checkout", getOrderHandler).Methods("GET")
	r.HandleFunc("/customer", CustomerInfoHandler).Methods("POST")
	r.HandleFunc("/tester", getCustomerHandler).Methods("GET")
	r.HandleFunc("/checkout", DeleteMealOrderHandler).Methods("POST")

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
	connString := "host=Foodie-db user=postgres password=12345 dbname=Foodie sslmode=disable"
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
