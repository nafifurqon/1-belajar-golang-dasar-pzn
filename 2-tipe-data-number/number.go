package main

import "fmt"

func main() {
	fmt.Println("Satu = ", 1) // default data type is int
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5) // default data type is float

	var test int8 = 100
	fmt.Println("Test = ", test)
	// test = 128 => cannot use 128 (untyped int constant) as int8 value in assignment (overflows)

	// byte = uint8
	// rune = int32
	// int = int32
	// uint = uint32
}
