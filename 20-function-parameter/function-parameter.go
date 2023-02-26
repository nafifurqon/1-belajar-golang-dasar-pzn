package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello,", firstName, lastName)
}

func main() {
	firstName := "Nafi"

	sayHelloTo(firstName, "Furqon")
	sayHelloTo("Muhammad", "Diani")
}
