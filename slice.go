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

	// Length and Capacity
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Afif"
	newSlice[1] = "Tri"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice)) //length
	fmt.Println(cap(newSlice)) //capacity

	newSlice2 := append(newSlice, "Rina")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Udin"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// Copy Slice
	fromSlice := days[:3]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Membedakan Array dan Slice
	iniArray := [...]int{1, 2, 3, 4}
	iniSlice := []int{1, 2, 3, 4}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
