package main

import "fmt"

func main() {
	//Operasi Matematika Dasar
	var a = 10
	var b = 20
	var c = 5
	var d = 5
	var e = (a * c) - 2*(b-d)

	fmt.Println(e)

	//	Augmented Assignment
	var i = 10
	i += 10

	fmt.Println(i)

	i += 5

	fmt.Println(i)

	//	unary operator
	var j = 3
	j++

	fmt.Println(j)

	j--
	fmt.Println(j)
}
