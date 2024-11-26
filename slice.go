package main

import "fmt"

func main() {
	names := [...]string{
		"Muhammad", "Afifudin", "Tri", "Setyorini", "Ifrikatul", "Khulda",
	}

	slice1 := names[4:6]

	fmt.Println(slice1)

	slice2 := names[2:]

	fmt.Println(slice2)

	slice3 := names[:4]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	var slice5 []string = names[:]
	fmt.Println(slice5)

	var long = len(slice5)
	fmt.Println(long)

	days := [...]string{
		"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	}

	var daySlice1 = days[3:]
	fmt.Println(daySlice1)

	daySlice1[0] = "kamis baru"
	daySlice1[1] = "jumat baru"

	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Coblosan")

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)
}
