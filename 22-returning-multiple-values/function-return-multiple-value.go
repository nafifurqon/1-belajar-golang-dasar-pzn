package main

import "fmt"

func getFullName() (string, string, string) {
	return "Nafi", "Furqon", "Diani"
}

func main() {
	// firstName, middleName, lastName := getFullName()
	firstName, middleName, _ := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	// fmt.Println(lastName)
}
