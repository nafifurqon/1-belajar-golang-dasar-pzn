package main

import "fmt"

func main() {
	name := "Nafi"
	counter := 0

	increment := func() {
		name := "Furqon"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
