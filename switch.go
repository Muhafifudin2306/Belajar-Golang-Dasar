package main

import "fmt"

func main() {
	name := "Afif"

	switch name {
	case "Afif":
		fmt.Println("Helo Afif")
	case "Afifudin":
		fmt.Println("Helo Afifudin")
	default:
		fmt.Println("Lets Connect")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama too long")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
