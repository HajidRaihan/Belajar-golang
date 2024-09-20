package main

import "fmt"

func main() {
	var name string = "pantek"
	fmt.Println(name)

	type NoKtp string

	const nomorKtp NoKtp = "lakjdalksjd"

	fmt.Println(nomorKtp)
}
