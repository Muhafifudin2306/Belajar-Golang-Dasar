package main

import "fmt"

func main() {
	// alias variable menggantikan string
	type NoKTP string

	var ktpEko NoKTP = "11111111111111"

	var contoh string = "222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(contohKtp)
	fmt.Println(ktpEko)
}
