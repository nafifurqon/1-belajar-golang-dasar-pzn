package main

import "fmt"

func main() {
	var name1 = "Nafi"
	var name2 = "Nafi"

	var result = name1 == name2
	fmt.Println(result)

	var name3 = "Nafi"
	var name4 = "Furqon"
	var result2 = name3 > name4
	fmt.Println(result2)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
