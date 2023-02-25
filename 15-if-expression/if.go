package main

import "fmt"

func main() {
	name := "Nafi"
	// name := "Furqon"
	// name := "Diani"

	if name == "Nafi" {
		fmt.Println("Hello Nafi")
	} else if name == "Furqon" {
		fmt.Println("Hello Furqon")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// length := len(name)
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
