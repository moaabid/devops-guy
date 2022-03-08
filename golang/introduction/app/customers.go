package main

type (
	Customer struct {
		FirstName string
		LastName  string
		Age       int
	}
)

func GetCustomersData() (customers []Customer) {

	mohammedaabid := Customer{FirstName: "Mohammed", LastName: "Aabid", Age: 30}

	customers = append(customers, mohammedaabid)

	return customers

}
