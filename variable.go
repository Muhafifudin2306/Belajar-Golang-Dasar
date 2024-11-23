package main

import "fmt"

func main(){
	// dengan tipe data
	var name string

	name  = "Muhammad Afifudin"
	
	fmt.Println("Hai nama saya ", name)

	// Tanpa tipe data

	var hobby = "Jogging"

	fmt.Println("Aku punya hobi ", hobby)

	// alternatif deklarasi variable (untuk pertama kali dan tidak dibuat ulang)
	food := "Ayam"

	fmt.Println("Saya suka makan", food)
	
	// Multiple varibale
	var (
		firstname = "Muhammad"
		lastname = "Afifudin"
	)
	fmt.Println("Nama saya ", firstname, lastname)
}