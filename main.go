package main

import (
	"html/template"
	// "io/ioutil"
	"log"
	"net/http"
	// "regexp"
)

func frontpageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/homepage.html"))
	tmpl.Execute(w, tmpl)
}
func orderHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/order.html"))
	tmpl.Execute(w, tmpl)
}

func main() {
	// http.HandleFunc("/view/", makeHandler(viewHandler))
	// http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/", frontpageHandler)   //home page
	http.HandleFunc("/order/", orderHandler) //order page

	//serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
