package main

import "fmt"

func main() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Print("Tiga koma lima = ", 3.5)

	var angka int16 = 3333
	var angka64 int64 = int64(angka)
	var angka32 int32 = int32(angka)

	fmt.Println("\n")
	fmt.Println(angka)
	fmt.Println(angka64)
	fmt.Println(angka32)
}