package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Muhammad"
	name[1] = "Nafi"
	name[2] = "Furqon"
	// name[3] = "Diani" // error: out of bounds

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{
		90,
		95,
		80,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(name))
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(len(lagi))

	// error
	// unlimitedArray := []int{}
	// unlimitedArray[0] = 1
	// fmt.Println(len(unlimitedArray))
}
