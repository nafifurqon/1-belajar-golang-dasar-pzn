package main

import "fmt"

func getGoodbye(name string) string {
	return "Good Bye " + name
}

func main() {
	sayGoodBye := getGoodbye
	result := sayGoodBye("Nafi")
	fmt.Println(result)
	fmt.Println(sayGoodBye("Nafi"))
}
