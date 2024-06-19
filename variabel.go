package main

import "fmt"

func main() {
	var name string

	name = "Alam Asy'arie"
	fmt.Println(name)

	name = "Alam"
	fmt.Println(name)

	// Deklarasi multi variabel
	var (
		firstName = "Alam"
		lastName  = "Asy'arie"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
