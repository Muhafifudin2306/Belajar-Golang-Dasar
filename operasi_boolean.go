package main

import "fmt"

func main(){
	var nilaiAbsen = 90
	var nilaiTugas  = 80

	var lulusAbsen bool = nilaiAbsen > 80
	var lulusTugas bool = nilaiTugas > 80

	var lulus bool = lulusAbsen && lulusTugas

	fmt.Println(lulus)
}