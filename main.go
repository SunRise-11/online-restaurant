package main

import (
	"html/template"
	// "io/ioutil"
	"log"
	"net/http"

	// "regexp"
	"github.com/gorilla/mux"
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
	r.HandleFunc("/", frontpageHandler).Methods("GET")
	r.HandleFunc("/order/", orderHandler).Methods("GET")

	//serve static files
	staticFileDirectory := http.Dir("./static")
	fs := http.FileServer(staticFileDirectory)
	staticFileHandler := http.StripPrefix("/static/", fs)
	http.Handle("/static/", staticFileHandler)
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	return r
}

func main() {

	// Declare a new router
	r := newRouter()

	// http.HandleFunc("/view/", makeHandler(viewHandler))
	// http.HandleFunc("/edit/", makeHandler(editHandler))
	// http.HandleFunc("/", frontpageHandler)   //home page
	// http.HandleFunc("/order/", orderHandler) //order page

	log.Fatal(http.ListenAndServe(":8080", r))
}
