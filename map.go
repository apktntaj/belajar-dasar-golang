package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Alam Asyarie",
		"address": "Indonesia",
	}

	person["title"] = "Programmer"

	fmt.Println(person)

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["istri"])

}