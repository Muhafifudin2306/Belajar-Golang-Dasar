package main

import "fmt"

func main() {
	counter := 1

	for counter < 10 {
		fmt.Println("Perulangan ke-", counter)
		counter++
	}

	//	bentuk lain
	for count := 1; count <= 10; count++ {
		fmt.Println("Bentuk 2 Perulangan ke-", count)
	}
	fmt.Println("Operation Finish!")

	//	Array
	name := []string{"Muhammad", "Afif", "Udin"}
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	for index, names := range name {
		fmt.Println("index", "=", index, "\n", "names", names)
	}

	for _, namer := range name {
		fmt.Println(namer)
	}
}
