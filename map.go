package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "hajid",
		"address": "barru",
	}

	delete(person, "address")
	fmt.Println(person)

	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}

	names := []string{"hajid", "agus", "pika"}

	for index, name := range names {
		fmt.Println("index", index, "name", name)
	}
}
