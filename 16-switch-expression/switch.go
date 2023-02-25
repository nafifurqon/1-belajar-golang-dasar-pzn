package main

import "fmt"

func main() {
	// name := "Nafi"
	name := "Furqon"
	// name := "Diani"

	switch name {
	case "Nafi":
		fmt.Println("Hello Nafi")
		fmt.Println("Hello Nafi")
	case "Furqon":
		fmt.Println("Hello Furqon")
		fmt.Println("Hello Furqon")
	default:
		fmt.Println("Kenalan donk")
		fmt.Println("Kenalan donk")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
