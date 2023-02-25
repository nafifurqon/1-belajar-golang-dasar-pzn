package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPNafi NoKTP = "123456787654321"
	var marriedStatus Married = true

	fmt.Println(noKTPNafi)
	fmt.Println(marriedStatus)
}
