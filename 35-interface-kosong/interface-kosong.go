package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	// var data int = Ups(1) // error
	var data interface{} = Ups(1)
	fmt.Println(data)
	data = Ups(2)
	fmt.Println(data)
	data = Ups(3)
	fmt.Println(data)
}
