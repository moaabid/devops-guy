package main

import (
	"fmt"
)

var test = "blah!"

func main() {
	customer := getCustomerData(1)
	fmt.Println(customer)
	customer = getCustomerData(3)
	fmt.Println(customer)
	fmt.Println(test)
	// customers := getCustomersData()
	// fmt.Println(customers)
	// fmt.Println(len(customers))

	customers := GetCustomersData()

	for _, customer := range customers {
		fmt.Println(customer)
	}

}

func getCustomerData(customerId int) (customer string) {

	if customerId == 1 {
		return "John"
	} else if customerId == 2 {
		return "Mohammed Aabid"
	} else {
		return "Unknown"
	}

}

// func getCustomersData() (customers []string) {
// 	customers = []string{"Mohammed Aabid", "John", "Ali"}

// 	customers = append(customers, "Peter")
// 	customers = append(customers, "Parker")

// 	// for x := 0; x < len(customers); x++ {
// 	// 	fmt.Println(customers[x])
// 	// }

// 	for _, customer := range customers {
// 		fmt.Println(customer)
// 	}

// 	return customers
// }
