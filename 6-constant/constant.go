package main

import "fmt"

func main() {
	const firstName string = "Nafi"
	const lastName = "Furqon"
	const value = 1000

	// Error: cannot assign to firstName (constant "Nafi" of type string)
	// firstName = "Budi"

	// Error: missing constant value
	// const firstName1 string
	// firstName1 = "Nafi"

	// fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	// firstName1 := "Muhammad" // this is variable, not constant

	// multiple constant declaration
	const (
		firstName1 = "Muhammad"
		lastName1  = "Nafi"
		value1     = 2000
	)
	fmt.Println(firstName1)
}
