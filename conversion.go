package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 32767
		nilai64 int64 = int64(nilai32)
		nilai16 int16 = int16(nilai32)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// Konversi dari string ke uint8 dan sebaliknya
	var name = "Afifudin"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
