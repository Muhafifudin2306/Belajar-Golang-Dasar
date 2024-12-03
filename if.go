package main

import "fmt"

func main() {
	name := "Afifudin"

	if name == "Afif" {
		fmt.Println("Hello Afif!")
	} else if name == "Afifudin" {
		fmt.Println("Hello Afifudin!")
	} else {
		fmt.Println("Hello World!")
	}

	//short statement
	if karakter := len(name); karakter < 5 {
		fmt.Println("karakter is less than 5")
	} else {
		fmt.Println("karakter is greater than 5")
	}
}
