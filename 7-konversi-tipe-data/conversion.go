package main

import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) // berubah value-nya karena terjadi integer overflow

	// integer overflow adalah
	// ketika sudah mencapai titik maksimum,
	// dia akan kembali ke titik awal (minimum)
	// ketika mencapai titik maksimum lagi,
	// dia kembali ke titik awal lagi

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Nafi"
	var e byte = name[0] // byte = uint8
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
