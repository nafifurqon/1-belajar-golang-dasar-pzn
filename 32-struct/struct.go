package main

import "fmt"

type Customer struct {
	// Name    string
	Name, Address string
	Age           int
	// Married       bool
}

func main() {
	var nafi Customer
	nafi.Name = "Nafi"
	nafi.Address = "Indonesia"
	nafi.Age = 25

	// fmt.Println(nafi)
	fmt.Println(nafi.Name)
	fmt.Println(nafi.Address)
	fmt.Println(nafi.Age)

	furqon := Customer{
		Name:    "Furqon",
		Address: "Indonesia",
		Age:     25,
	}
	fmt.Println(furqon)

	diani := Customer{"Diani", "Indonesia", 25} // error ketika jumlah fields tidak sesuai
	fmt.Println(diani)
}
