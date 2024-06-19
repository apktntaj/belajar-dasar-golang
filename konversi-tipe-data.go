package main

import "fmt"

func main() {
	nilai32 := 1000000
	nilai64 := int64(nilai32)
	nilai16 := int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16) // hasil akan berubah karena nilai32 tidak bisa di konversi ke int16
}