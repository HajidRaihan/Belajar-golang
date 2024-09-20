package main

import "fmt"

func main() {
	var values = [10]int{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}

	slice1 := values[3:5]

	days := [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	daySlice := days[5:]

	var newSlice []string = make([]string, 2, 5)

	newSlice[0] = "hajid"
	newSlice[1] = "raihan"

	newSlice = append(newSlice, "pantek")

	fmt.Println(daySlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(values)
	fmt.Println(slice1)
	fmt.Println(newSlice)
	fmt.Println(cap(iniArray))
	fmt.Println(cap(iniSlice))
}
