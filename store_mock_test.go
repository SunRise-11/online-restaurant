package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetMealsHandler(t *testing.T) {
	// Initialize the mock store
	mockStore := InitMockStore()

	/* Define the data that we want to return when the mocks `GetMeals` method is
	called
	Also, we expect it to be called only once
	*/
	mockStore.On("GetMeals").Return([]*Meal{
		{"Meat Pizza", 20, "food1.png"},
	}, nil).Once()

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(mockGetMealHandler)

	// Now, when the handler is called, it should call our mock store, instead of
	// the actual one
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Meal{"Meat Pizza", 20, "food1.png"}
	m := []Meal{}
	err = json.NewDecoder(recorder.Body).Decode(&m)

	if err != nil {
		t.Fatal(err)
	}

	actual := m[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

	// the expectations that we defined in the `On` method are asserted here
	mockStore.AssertExpectations(t)
}

func TestGetOrdersHandler(t *testing.T) {
	// Initialize the mock store
	mockStore := InitMockStore()

	/* Define the data that we want to return when the mocks `GetOrders` method is
	called
	Also, we expect it to be called only once
	*/
	mockStore.On("GetOrders").Return([]*Order{
		{1, "Spaghetti", 10, "spag.jpg", 5, 50},
	}, nil).Once()

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(mockGetOrderHandler)

	// Now, when the handler is called, it should call our mock store, instead of
	// the actual one
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Order{1, "Spaghetti", 10, "spag.jpg", 5, 50}
	m := []Order{}
	err = json.NewDecoder(recorder.Body).Decode(&m)

	if err != nil {
		t.Fatal(err)
	}

	actual := m[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

	// the expectations that we defined in the `On` method are asserted here
	mockStore.AssertExpectations(t)
}

func TestCreateOrderHandler(t *testing.T) {

	mockStore := InitMockStore()
	/*
	 Similarly, we define our expectations for the `CreateOrder` method.
	 We expect the first argument to the method to be the order struct
	 defined below, and tell the mock to return a `nil` error
	*/
	mockStore.On("CreateOrders", &Order{0, "meatpizza", 5, "pizza.png", 5, 25}).Return(nil)

	form := newCreateOrderForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(createOrderHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)
}

func TestCreateCustomerInfoHandler(t *testing.T) {

	mockStore := InitMockStore()
	/*
	 Similarly, we define our expectations for the `CreateCustomers` method.
	 We expect the first argument to the method to be the customers struct
	 defined below, and tell the mock to return a `nil` error
	*/
	mockStore.On("DeleteOrders").Return(nil)
	mockStore.On("CreateCustomer", &Customer{"Faruq", "Hidden Leaf", "2 plates of Spaghetti", 22}).Return(nil)

	form := newCreateCustomerForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(CustomerInfoHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)
}

func newCreateOrderForm() *url.Values {
	form := url.Values{}
	form.Set("meal", "meatpizza")
	form.Set("price", "5")
	form.Set("image", "pizza.png")
	form.Set("PlateNumbers", "5")
	return &form
}

func newCreateCustomerForm() *url.Values {
	form := url.Values{}
	form.Set("customer", "Faruq")
	form.Set("address", "Hidden Leaf")
	form.Set("meal", "2 plates of Spaghetti")
	form.Set("total", "22")
	return &form
}
