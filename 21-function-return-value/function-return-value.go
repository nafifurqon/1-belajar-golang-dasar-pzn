package main

import "fmt"

func getHello(name string) string {
	// return "Hello, " + name

	if name == "" {
		return "Hello, bro!"
	} else {
		return "Hello, " + name
	}
}

func main() {
	name := "Nafi"
	result := getHello(name)
	fmt.Println(result)

	fmt.Println(getHello(""))
	// getHello("Furqon")
}
