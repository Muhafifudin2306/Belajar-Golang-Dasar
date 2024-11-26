package main

import "fmt"

func main() {
	var nama [3]string

	nama[0] = "Muhammad"
	nama[1] = "Afifudin"
	nama[2] = "Udin"

	fmt.Println(nama)

	fmt.Println(nama[0])
	fmt.Println(nama[2])
	fmt.Println(nama[1])

	var values = [...]int{
		10,
		20,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(len(values))

	values[0] = 1

	fmt.Println(values[0])
}
