package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Meal struct {
	Food  string `json:"foodName"`
	Price int    `json:"cost"`
	Image string `json:"imagelink"`
}
type Order struct {
	ID     int    `json:"id"`
	Meal   string `json:"foodName"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
	Plates int    `json:"plateNumbers"`
}

type Customer struct {
	Name    string `json:"customer"`
	Address string `json:"address"`
	Plates  int    `json:"plateNumbers"`
	Meal    string `json:"foodName"`
	Price   int    `json:"price"`
}

func getMealHandler(w http.ResponseWriter, r *http.Request) {

	meal, err := store.GetMeals()
	if err != nil {
		fmt.Println(err)
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
	order.Meal = r.Form.Get("meal")
	order.Price, err = strconv.Atoi(r.Form.Get("price"))
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	order.Image = r.Form.Get("image")
	order.Plates, err = strconv.Atoi(r.Form.Get("PlateNumbers"))
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = store.CreateOrders(&order)
	if err != nil {
		fmt.Println(err)
	}

	//Finally, we redirect the user to the original HTMl page
	// (located at `/static/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/static/order.html", http.StatusFound)
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {

	order, err := store.GetOrders()
	if err != nil {
		fmt.Println(err)
	}

	tmpl := template.Must(template.New("").ParseFiles("./static/checkout.html"))

	context := struct{ Orders []*Order }{order}

	err = tmpl.ExecuteTemplate(w, "checkout.html", context)

	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
func getCustomerHandler(w http.ResponseWriter, r *http.Request) {

	customer, err := store.GetCustomer()
	if err != nil {
		fmt.Println(err)
	}

	tmpl := template.Must(template.New("").ParseFiles("./static/tester.html"))

	context := struct{ Customers []*Customer }{customer}

	err = tmpl.ExecuteTemplate(w, "tester.html", context)

	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
	customer.Name = r.Form.Get("customer")
	customer.Address = r.Form.Get("address")
	customer.Meal = r.Form.Get("meal")
	customer.Price, err = strconv.Atoi(r.Form.Get("price"))
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	customer.Plates, err = strconv.Atoi(r.Form.Get("plates"))
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = store.CreateCustomer(&customer)
	if err != nil {
		fmt.Println(err)
	}
	err1 := store.DeleteOrders()
	if err1 != nil {
		fmt.Println(err1)
	}

	//Finally, we redirect the user to the original HTMl page
	// (located at `/static/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/static/", http.StatusFound)
}

func DeleteMealOrderHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var id int

	// Get the information about the meal to be deleted from the form info
	id, err = strconv.Atoi(r.Form.Get("delete"))

	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err1 := store.DeleteMealOrder(int64(id))
	if err1 != nil {
		fmt.Println(err1)
	}

	order, err := store.GetOrders()
	if err != nil {
		fmt.Println(err)
	}

	tmpl := template.Must(template.New("").ParseFiles("./static/checkout.html"))

	context := struct{ Orders []*Order }{order}

	err = tmpl.ExecuteTemplate(w, "checkout.html", context)

	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

// mockstore handlers
func mockGetMealHandler(w http.ResponseWriter, r *http.Request) {

	meal, err := store.GetMeals()
	if err != nil {
		fmt.Println(err)
	}
	// Convert to json data for mock test
	MealBytes, err := json.Marshal(meal)
	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(MealBytes)
}

func mockGetOrderHandler(w http.ResponseWriter, r *http.Request) {

	order, err := store.GetOrders()
	if err != nil {
		fmt.Println(err)
	}

	// Convert to json data for mock test
	OrderBytes, err := json.Marshal(order)
	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(OrderBytes)

}
