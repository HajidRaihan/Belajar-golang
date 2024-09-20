package main

import "fmt"

func main() {
	var values = [3]int{
		1, 2, 3,
	}

	var buah [3]string

	// tidak menentukan panjang array
	var sembarang = [...]int{
		1, 2, 3,
	}
	buah[0] = "apel"
	buah[1] = "pisang"
	buah[2] = "jeruk"
	fmt.Println(values)
	fmt.Println(buah)
	fmt.Println(sembarang)
}
