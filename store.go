package main

// The sql go library is needed to interact with the database
import (
	"database/sql"
	"fmt"
)

// Our store will have methods to create and get all  meals, customers, orders and delete orders
// Each method returns an error, in case something goes wrong
type Store interface {
	CreateCustomer(customer *Customer) error
	CreateOrders(order *Order) error
	GetMeals() ([]*Meal, error)
	GetOrders() ([]*Order, error)
	GetCustomer() ([]*Customer, error)
	DeleteOrders() error
}

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection.
type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateMeals(meal *Meal) error {
	// 'Meal' is a simple struct which has "Food", "Price" and "Image Link" attributes
	// THe first underscore means that we don't care about what's returned from
	// this insert query. We just want to know if it was inserted correctly,
	// and the error will be populated if it wasn't
	_, err := store.db.Query("INSERT INTO meal(food, price, image) VALUES ($1,$2,$3)", meal.Food, meal.Price, meal.Image)
	return err
}

func (store *dbStore) CreateOrders(orders *Order) error {
	// 'Meal' is a simple struct which has "foodname" attributes
	// THe first underscore means that we don't care about what's returned from
	// this insert query. We just want to know if it was inserted correctly,
	// and the error will be populated if it wasn't

	var id int64
	err := store.db.QueryRow("INSERT INTO orders(meal,price,image,plates) VALUES ($1,$2,$3,$4)RETURNING id", orders.Meal, orders.Price, orders.Image, orders.Plates).Scan(&id)

	if err != nil {
		fmt.Println(fmt.Errorf("unable to execute the query: %v", err))
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return err
}

func (store *dbStore) CreateCustomer(customer *Customer) error {
	// 'Customer' is a simple struct which has "Name","Address","Plates","Meal" and "Price" attributes

	_, err := store.db.Query("INSERT INTO customers(customer_name,location_address,plates,meal,price) VALUES ($1,$2,$3,$4,$5)", customer.Name, customer.Address, customer.Plates, customer.Meal, customer.Price)
	return err

}
func (store *dbStore) DeleteOrders() error {
	// Delete meal orders after checkout

	_, err1 := store.db.Query("DELETE FROM orders")
	return err1
}

func (store *dbStore) GetMeals() ([]*Meal, error) {
	// Query the database for every meal, and return the result to the
	// `rows` object
	rows, err := store.db.Query("SELECT food, price, image from meal")
	// We return incase of an error, and defer the closing of the row structure
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create the data structure that is returned from the function.
	// By default, this will be an empty array of meals
	meals := []*Meal{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a meal,
		meal := &Meal{}
		// Populate the `Food`, `Price` and `Image` attributes of the meal,
		// and return incase of an error
		if err := rows.Scan(&meal.Food, &meal.Price, &meal.Image); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		meals = append(meals, meal)
	}
	return meals, nil
}
func (store *dbStore) GetOrders() ([]*Order, error) {
	// Query the database for the order, and return the result to the
	// `rows` object
	rows, err := store.db.Query("SELECT id,meal,price,image,plates from orders")

	// We return incase of an error, and defer the closing of the row structure
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create the data structure that is returned from the function.
	// By default, this will be an empty array of meals
	orders := []*Order{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to an order,
		order := &Order{}
		// Populate the `ordered meal` and `number of plates` attributes of the meal,
		// and return incase of an error
		if err := rows.Scan(&order.ID, &order.Meal, &order.Price, &order.Image, &order.Plates); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		orders = append(orders, order)
	}
	return orders, nil
}
func (store *dbStore) GetCustomer() ([]*Customer, error) {
	// Query the database for the order, and return the result to the
	// `rows` object
	rows, err := store.db.Query("SELECT customer_name, location_address,plates,meal,price from customers")
	// We return incase of an error, and defer the closing of the row structure
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create the data structure that is returned from the function.
	// By default, this will be an empty array of meals
	customers := []*Customer{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to an order,
		customer := &Customer{}
		// Populate the `ordered meal` and `number of plates` attributes of the meal,
		// and return incase of an error
		if err := rows.Scan(&customer.Name, &customer.Address, &customer.Plates, &customer.Meal, &customer.Price); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		customers = append(customers, customer)
	}
	return customers, nil
}

// The store variable is a package level variable that will be available for
// use throughout our application code
var store Store

/*
We will need to call the InitStore method to initialize the store. This will
typically be done at the beginning of our application (in this case, when the server starts up)
This can also be used to set up the store as a mock, which we will be observing
later on
*/
func InitStore(s Store) {
	store = s
}
