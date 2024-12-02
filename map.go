package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Muhammad Afifudin",
		"address": "Central Java",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Happy Money"
	book["author"] = "Afif"
	book["error"] = "Error"

	fmt.Println(book)

	delete(book, "error")

	fmt.Println(book)
}
