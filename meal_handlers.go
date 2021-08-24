package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Meal struct {
	Food  string `json:"foodName"`
	Price int    `json:"cost"`
	Image string `json:"imagelink"`
}

func getMealHandler(w http.ResponseWriter, r *http.Request) {

	meal, err := store.GetMeals()
	if err != nil {
		fmt.Println(err)
	}

	//Convert the "meals" variable to json
	foodListBytes, err := json.Marshal(meal)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of meals to the response
	w.Write(foodListBytes)
}
