package main

import (
	"database/sql"
	"testing"

	// The "testify/suite" package is used to make the test suite
	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite
	/*
		The suite is defined as a struct, with the store and db as its
		attributes. Any variables that are to be shared between tests in a
		suite should be stored as attributes of the suite instance
	*/
	store *dbStore
	db    *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	/*
		The database connection is opened in the setup, and
		stored as an instance variable,
		as is the higher level `store`, that wraps the `db`
	*/
	connString := "host=Foodie-db user=postgres password=12345 dbname=Foodie sslmode=disable"
	// connString := "user=postgres password=12345 dbname=Foodie sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuite) SetupTest() {
	/*
		We delete all entries from the tables before each test runs, to ensure a
		consistent state before our tests run. In more complex applications, this
		is sometimes achieved in the form of migrations
	*/
	_, err := s.db.Query("DELETE FROM meal")
	_, err1 := s.db.Query("DELETE FROM orders")
	_, err2 := s.db.Query("DELETE FROM customers")
	if (err != nil) && (err1 != nil) && (err2 != nil) {
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite() {
	// Close the connection after all tests in the suite finish
	s.db.Close()
}

// This is the actual "test" as seen by Go, which runs the tests defined below
func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestCreateMeal() {
	// Create a meal through the store `CreateMeal` method
	s.store.CreateMeals(&Meal{
		Food:  "foodName",
		Price: 10,
		Image: "imagelink",
	})

	// Query the database for the entry we just created
	res, err := s.db.Query(`SELECT COUNT(*) FROM meal WHERE Food='foodName' AND Price=10 AND Image= 'imagelink' `)
	if err != nil {
		s.T().Fatal(err)
	}

	// Get the count result
	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}

	// Assert that there must be one entry with the properties of the meal that we just inserted (since the database was empty before this)
	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
}
func (s *StoreSuite) TestCreateOrders() {
	// Create an order through the store `CreateOrder` method
	s.store.CreateOrders(&Order{
		Meal:      "Jollof Rice",
		Price:     10,
		Image:     "JollofRice.jpeg",
		Plates:    5,
		TotalCost: 50,
	})

	// Query the database for the entry we just created
	res, err := s.db.Query(`SELECT COUNT(*) FROM orders WHERE Meal='Jollof Rice' AND Price=10 AND Image='JollofRice.jpeg' AND Plates=5 AND TotalCost=50`)

	if err != nil {
		s.T().Fatal(err)
	}

	// Get the count result
	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}

	// Assert that there must be one entry with the properties of the meal that we just inserted (since the database was empty before this)
	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
}
func (s *StoreSuite) TestCreateCustomer() {
	// Create an order through the store `CreateCustomer` method
	s.store.CreateCustomer(&Customer{
		Name:      "customer",
		Address:   "address",
		Meal:      "5 plates of spaghetti",
		TotalCost: 22,
	})

	// Query the database for the entry we just created
	res, err := s.db.Query(`SELECT COUNT(*) FROM customers WHERE customer_name='customer' AND location_address='address' AND meal= '5 plates of spaghetti' AND TotalCost=22 `)
	if err != nil {
		s.T().Fatal(err)
	}

	// Get the count result
	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}

	// Assert that there must be one entry with the properties of the meal that we just inserted (since the database was empty before this)
	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
}

func (s *StoreSuite) TestGetMeal() {
	// Insert a sample meal into the `meal` table
	_, err := s.db.Query(`INSERT INTO meal (Food, Price, Image) VALUES('Yam',10, 'yam.jpg')`)
	if err != nil {
		s.T().Fatal(err)
	}

	// Get the list of meals through the stores `GetMeals` method
	meals, err := s.store.GetMeals()
	if err != nil {
		s.T().Fatal(err)
	}

	// Assert that the count of meals received must be 1
	nMeals := len(meals)
	if nMeals != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", nMeals)
	}

	// Assert that the details of the meal is the same as the one we inserted
	expectedMeal := Meal{"Yam", 10, "yam.jpg"}
	if *meals[0] != expectedMeal {
		s.T().Errorf("incorrect details, expected %v, got %v", expectedMeal, *meals[0])
	}
}
func (s *StoreSuite) TestGetOrders() {
	// Insert a sample order into the `orders` table
	_, err := s.db.Query(`INSERT INTO orders (id,meal, price,image, plates, totalcost) VALUES(7, 'Egg Rice', 5, 'eggrice.jpg', 5, 25)`)

	if err != nil {
		s.T().Fatal(err)
	}

	// Get the list of orders through the stores `GetOrders` method
	order, err := s.store.GetOrders()
	if err != nil {
		s.T().Fatal(err)
	}

	// Assert that the count of order received must be 1
	nOrders := len(order)
	if nOrders != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", nOrders)
	}

	// Assert that the details of the order is the same as the one we inserted
	expectedOrder := Order{7, "Egg Rice", 5, "eggrice.jpg", 5, 25}
	if *order[0] != expectedOrder {
		s.T().Errorf("incorrect details, expected %v, got %v", expectedOrder, *order[0])
	}
}
