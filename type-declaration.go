package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTP NoKTP = "7371012601880005"
	var marriedStatus Married = true


	fmt.Println(noKTP)
	fmt.Println(marriedStatus)
}