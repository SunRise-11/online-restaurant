package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Meal struct {
	Food  string `json:"foodName"`
	Price int    `json:"cost"`
	Image string `json:"imagelink"`
}
type Order struct {
	OrderedMeal string `json:"foodName"`
}
type Customer struct {
	Name    string `json:"customer"`
	Address string `json:"address"`
}

func getMealHandler(w http.ResponseWriter, r *http.Request) {

	meal, err := store.GetMeals()
	if err != nil {
		fmt.Println(err)
	}

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// tmpl := template.New("./static/homepage.html")
	tmpl := template.Must(template.New("").ParseFiles("./static/homepage.html"))
	// tmpl := template.Must(template.ParseFiles("./static/homepage.html"))
	context := struct{ Meals []*Meal }{meal}
	// tmpl.Execute(w, tmpl, context)
	err = tmpl.ExecuteTemplate(w, "homepage.html", context)

	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	order := Order{}

	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the ordered meal from the form info
	order.OrderedMeal = r.Form.Get("foodName")
	// order.Image = r.Form.Get("imagelink")

	err = store.CreateOrders(&order)
	if err != nil {
		fmt.Println(err)
	}

	//Finally, we redirect the user to the original HTMl page
	// (located at `/static/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/static/", http.StatusFound)
}

func CustomerInfoHandler(w http.ResponseWriter, r *http.Request) {
	customer := Customer{}

	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the customer from the form info
	customer.Name = r.Form.Get("user")
	customer.Address = r.Form.Get("address")

	err = store.CreateCustomer(&customer)
	if err != nil {
		fmt.Println(err)
	}

	//Finally, we redirect the user to the original HTMl page
	// (located at `/static/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/static/", http.StatusFound)
}
