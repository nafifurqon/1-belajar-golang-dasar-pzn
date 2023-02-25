package main

import "fmt"

func main() {
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	for counter := 0; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	fmt.Println("=========================================")

	slice := []string{"Muhammad", "Nafi", "Furqon", "Diani"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	fmt.Println("=========================================")
	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}
	fmt.Println("=========================================")
	for _, value := range slice {
		fmt.Println("value =", value)
	}
	fmt.Println("=========================================")

	person := make(map[string]string)
	person["name"] = "Nafi"
	person["title"] = "Programmer"
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
