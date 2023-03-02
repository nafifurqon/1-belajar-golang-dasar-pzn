package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Say Huuu from", a.Name)
}

func main() {
	var nafi Customer
	nafi.Name = "Nafi"
	nafi.Address = "Indonesia"
	nafi.Age = 25

	// sayHi(nafi, "Muhammad") // bukan struct method
	nafi.sayHi("Muhammad")
	nafi.sayHuuu()

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
