package main

import "fmt"

func main() {
	var name string = "Nafi"
	fmt.Println(name)

	name = "Nafi Furqon"
	fmt.Println(name)

	name = "Nafi Diani"
	fmt.Println(name)

	// name = true // error because different type

	var friendName = "Budi"
	fmt.Println(friendName)

	var age int8 = 25
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Muhammad"
		lastName  = "Nafi"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
