package main

import "fmt"

func main() {
	names:= [...]string {"Alam", "Aida", "Raya", "Masyitah"}
	slice:= names[0:3]
	slice2 := slice

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice[0] = "Alam Asyarie"
	slice2[1] = "Aida Raya Masyitah"
	fmt.Println(names)


}