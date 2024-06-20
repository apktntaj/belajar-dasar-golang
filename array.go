package main

import "fmt"

func main() {
	var names [4]string

	names[0] = "Alam"
	names[1] = "Aida"
	names[2] = "Raya"
	names[3] = "Masyitah"

	fmt.Println(names[0], names[1], names[2], names[3])

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values[0], values[1], values[2])

	var names2 [10]string

	names2[0] = "Alam"
	names2[1] = "Aida"
	names2[2] = "Raya"
	names2[3] = "Masyitah"

	fmt.Println(names2[0], names2[1], names2[2], names2[3])
	fmt.Println(len(names2))
}