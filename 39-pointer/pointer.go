package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	address2 := address1  // pass by value
	address3 := &address1 // pass by reference
	address4 := &address1

	address2.City = "Bogor"
	address3.City = "Bogor"

	// address3 = &Address{"Malang", "Jawa Timur", "Indonesia"}
	*address3 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println("address3", address3)
	fmt.Println("address1", address1) // address1 tidak berubah
	// fmt.Println(address2)
	fmt.Println("address3", address3)
	fmt.Println("address4", address4)

	var address5 *Address = new(Address)
	address5.City = "Jakarta"
	fmt.Println("address5", address5)
}
